package domain

import "github.com/aws/aws-lambda-go/events"

type RespAPI struct {
	Status 			int `json:"status"`
	Message 		string `json:"message"`
	CustomResp 	*events.APIGatewayProxyResponse `json:"customResp"`
}