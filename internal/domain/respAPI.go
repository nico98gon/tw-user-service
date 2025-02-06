package domain

import "github.com/aws/aws-lambda-go/events"

type RespAPI struct {
	Status 			int 														`json:"status"`
	Message 		string 													`json:"message"`
	Data    		interface{} 										`json:"data,omitempty"`
	Meta    		interface{} 										`json:"meta,omitempty"`
	CustomResp 	*events.APIGatewayProxyResponse `json:"customResp"`
}