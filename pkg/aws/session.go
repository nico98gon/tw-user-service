package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func StartAWS() {
	Ctx = context.TODO()
	// region := os.Getenv("AWS_REGION")

	if os.Getenv("APP_ENV") == "local" {
		// Modo local: Usar credenciales del archivo ~/.aws/credentials
		Cfg, err = config.LoadDefaultConfig(Ctx, config.WithSharedConfigProfile("default"))
		if err != nil {
			log.Fatalf("Error al cargar la configuración de AWS en local: %v", err)
		}
	} else {
		// Modo Lambda: Usar el rol de IAM asociado a la función Lambda
		Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
		if err != nil {
			log.Fatalf("Error al cargar la configuración de AWS en Lambda: %v", err)
		}
	}
}