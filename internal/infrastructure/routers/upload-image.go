package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"strings"
	"user-service/internal/domain"
	"user-service/internal/domain/users"
	"user-service/internal/infrastructure/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offsest int64, whence int) (int64, error) {
	return 0, nil
}

func isValidImageType(contentType string) (bool, string) {
	switch contentType {
	case "image/jpeg":
		return true, ".jpg"
	case "image/png":
		return true, ".png"
	default:
		return false, ""
	}
}

func isBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim *domain.Claim) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400

	IDUser := claim.ID.Hex()
	if len(IDUser) == 0 || IDUser == "000000000000000000000000" {
		r.Message = "Error: ID de usuario no válido en el token"
		fmt.Println(r.Message)
		return r
	}

	var fileName string
	var user users.User

	bucket := aws.String(ctx.Value(domain.Key("bucket_name")).(string))
	if bucket == nil {
		r.Message = "Error: No se pudo obtener el bucket de la solicitud"
		fmt.Println(r.Message)
		return r
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil || !strings.HasPrefix(mediaType, "multipart/") {
		r.Status = 400
		r.Message = "Error: el Content-Type debe ser 'multipart/form-data'"
		fmt.Println(r.Message)
		return r
	}

	if len(request.Body) == 0 {
		r.Status = 400
		r.Message = "Error: el cuerpo de la solicitud está vacío"
		fmt.Println(r.Message)
		return r
	}

	var body []byte
	if isBase64(request.Body) {
		body, err = base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			r.Status = 500
			r.Message = "Error al parsear el body: " + err.Error()
			fmt.Println(r.Message)
			return r
		}
	} else {
		body = []byte(request.Body)
	}

	mr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
	p, err := mr.NextPart()
	if err != nil && err != io.EOF {
		r.Status = 500
		r.Message = "Error al parsear el body: " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if err != io.EOF {
		if p.FileName() != "" {
			contentType := p.Header.Get("Content-Type")
			isValid, ext := isValidImageType(contentType)
			if !isValid {
				r.Status = 400
				r.Message = "Error: El archivo debe ser JPG o PNG"
				fmt.Println(r.Message)
				return r
			}

			switch uploadType {
			case "A":
				fileName = "avatars/" + IDUser + ext
				user.Avatar = fileName
			case "B":
				fileName = "banners/" + IDUser + ext
				user.Banner = fileName
			}

			buff := bytes.NewBuffer(nil)
			if _, err := io.Copy(buff, p); err != nil {
				r.Status = 500
				r.Message = "Error al parsear el body: " + err.Error()
				fmt.Println(r.Message)
				return r
			}

			sess, err := session.NewSession(&aws.Config{
				Region: aws.String("us-east-1")},
			)
			if err != nil {
				r.Status = 500
				r.Message = "Error al crear la sesión: " + err.Error()
				fmt.Println(r.Message)
				return r
			}

			uploader := s3manager.NewUploader(sess)
			_, err = uploader.Upload(&s3manager.UploadInput{
				Bucket: bucket,
				Key:    aws.String(fileName),
				Body:   &readSeeker{Reader: buff},
			})
			if err != nil {
				r.Status = 500
				r.Message = "Error al subir la imagen: " + err.Error()
				fmt.Println(r.Message)
				return r
			}
		}

		status, err := db.UpdateRegister(user, IDUser)
		if err != nil || !status {
			r.Status = 400
			r.Message = "Error al modificar el registro del usuario: " + err.Error()
			fmt.Println(r.Message)
			return r
		}
	}

	r.Status = 200
	r.Message = "Imagen subida correctamente"
	fmt.Println(r.Message)
	return r
}