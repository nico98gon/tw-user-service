package main

import (
	"fmt"
	"os"
	"user-service/server"
	"user-service/services"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "local" {
		fmt.Println("Corriendo en local...")
		server.StartLocalServer()
	} else {
		fmt.Println("Corriendo en AWS Lambda...")
		lambda.Start(services.LambdaExec)
	}
}