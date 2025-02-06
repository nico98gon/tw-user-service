package utils

import (
	"github.com/aws/aws-lambda-go/events"
)

func HandleError(status int, message string) *events.APIGatewayProxyResponse {
	return FormatResponse(status, message, nil, nil)
}