package aws

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func StartAWS() {
	Ctx = context.TODO()
	region := os.Getenv("AWS_REGION")
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion(region))
	if err != nil {
		panic("Error al cargar la configuración de AWS" + err.Error())
	}
}