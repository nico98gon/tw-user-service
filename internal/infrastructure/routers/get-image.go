package routers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"user-service/internal/domain"
	"user-service/internal/infrastructure/db"
	awsSession "user-service/pkg/aws"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) == 0 || ID == "000000000000000000000000" {
		r.Message = "ID es requerido"
		return r
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		r.Message = "Error al buscar el perfil: " + err.Error()
		return r
	}

	var fileName string
	switch uploadType {
	case "A":
		fileName = profile.Avatar
	case "B":
		fileName = profile.Banner
	}
	if fileName == "" {
		r.Status = 404
		r.Message = "No se encontró la imagen en el perfil del usuario"
		return r
	}

	svc := s3.NewFromConfig(awsSession.Cfg)

	file, err := downloadFromS3(ctx, svc, fileName)
	if err != nil {
		r.Status = 500
		r.Message = "Error al descargar el archivo de S3: " + err.Error()
		return r
	}

	r.CustomResp = &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/octet-stream",
			"Content-Disposition": fmt.Sprintf("attachment; filename=%s", fileName),
		},
		Body: file.String(),
	}

	return r
}

func downloadFromS3(ctx context.Context, svc *s3.Client, fileName string) (*bytes.Buffer, error) {
	bucket := ctx.Value(domain.Key("bucket_name")).(string)
	obj, err := svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return nil, err
	}
	defer obj.Body.Close()
	fmt.Println("bucketname: ", bucket)

	file, err := io.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(file)

	return buffer, nil
}