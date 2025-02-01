package cmd

import (
	"user-service/services"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(services.LambdaExec)
}

