package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"user-service/internal/domain"
	"user-service/internal/domain/users"
	"user-service/internal/infrastructure/db"
	jwt "user-service/pkg/JWT"

	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) domain.RespAPI {
	var u users.User
	var r domain.RespAPI
	r.Status = 400

	body := ctx.Value(domain.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		r.Message = "Error al parsear el body"+err.Error()
		fmt.Println(r.Message)
		return r
	}

	if err := users.LoginValidations(u); err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	userData, found := db.TryLogin(u.Email, u.Password)
	if !found {
		r.Message = "Usuario y/o contraseña incorrectos"
		fmt.Println(r.Message)
		return r
	}

	jwtKey, err := jwt.GenerateToken(ctx, userData)
	if err != nil {
		r.Message = "Error al intentar generar el token > "+err.Error()
		fmt.Println(r.Message)
		return r
	}

	resp := users.LoginResponse{
		Token: jwtKey,
	}

	token, err := json.Marshal(resp)
	if err != nil {
		r.Message = "Error al intentar generar el token > "+err.Error()
		fmt.Println(r.Message)
		return r
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    jwtKey,
		Expires:  time.Now().Add(time.Hour * 24),
	}
	cookieStr := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:	string(token),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":	cookieStr,
		},
	}

	r.Status = 200
	r.Message = "OK"
	r.CustomResp = res

	return r
}