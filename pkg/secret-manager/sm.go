package secretmanager

import (
	"encoding/json"
	"fmt"
	"os"
	domain "user-service/internal/domain/sm"
	awsSession "user-service/pkg/aws"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// GetSecret obtiene las claves secretas de AWS o de un archivo local
func GetSecret(secretName string) (domain.Secret, error) {
	var secretData domain.Secret

	useAWSSecrets := os.Getenv("USE_AWS_SECRETS") == "true"

	if os.Getenv("APP_ENV") == "local" && useAWSSecrets {
		fmt.Println("> [LOCAL] Usando AWS Secrets Manager")
		return getAWSSecret(secretName)
	}

	if os.Getenv("APP_ENV") == "local" {
		fmt.Println("> [LOCAL] Usando variables de entorno locales")
		secretData = domain.Secret{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			IsServ:   os.Getenv("DB_IS_SRV") == "true",
			Database: os.Getenv("DB_DATABASE"),
			JWTSign:  os.Getenv("JWT_SIGN"),
		}
		return secretData, nil
	}

	fmt.Println("> [AWS] Usando AWS Secrets Manager")
	return getAWSSecret(secretName)
}

// Función separada para obtener secretos de AWS
func getAWSSecret(secretName string) (domain.Secret, error) {
	var secretData domain.Secret

	svc := secretsmanager.NewFromConfig(awsSession.Cfg)
	key, err := svc.GetSecretValue(awsSession.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return secretData, fmt.Errorf("error al obtener secreto: %v", err)
	}

	err = json.Unmarshal([]byte(*key.SecretString), &secretData)
	if err != nil {
		return secretData, fmt.Errorf("error deserializando secreto: %v", err)
	}

	return secretData, nil
}