package handlers

import (
	"context"
	"fmt"
	"user-service/internal/domain"

	"github.com/aws/aws-lambda-go/events"
)

func AwsHandler(ctx context.Context, request events.APIGatewayProxyRequest) domain.RespAPI {
	// Envía a CloudWatch Logs
	fmt.Println("Procesando"+ctx.Value(domain.Key("path")).(string) + " > " + ctx.Value(domain.Key("method")).(string))

	var r domain.RespAPI
	r.Status = 400

	switch ctx.Value(domain.Key("method")).(string) {
	case "GET":
		switch ctx.Value(domain.Key("path")).(string) {
		
		}
		//
	case "POST":
		switch ctx.Value(domain.Key("path")).(string) {
		
		}
		//
	case "PUT":
		switch ctx.Value(domain.Key("path")).(string) {
		
		}
		//
	case "DELETE":
		switch ctx.Value(domain.Key("path")).(string) {
		
		}
		//
	}

	r.Message = "Method Invalid"
	return r
}