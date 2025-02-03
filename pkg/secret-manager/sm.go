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

	if os.Getenv("APP_ENV") == "local" {
		// En local: Leer secretos desde un archivo o variables de entorno
		secretData = domain.Secret{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Database: os.Getenv("DB_DATABASE"),
			JWTSign:  os.Getenv("JWT_SIGN"),
		}
		// fmt.Println("> Usando secretos locales:", secretData)
		return secretData, nil
	}

	// En Lambda: Obtener secretos desde AWS Secrets Manager
	fmt.Println("> Se pide secreto desde AWS Secrets Manager:", secretName)

	svc := secretsmanager.NewFromConfig(awsSession.Cfg)
	key, err := svc.GetSecretValue(awsSession.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println("Error al obtener el secreto:", err.Error())
		return secretData, err
	}

	fmt.Println("Secreto obtenido:", *key.SecretString)

	err = json.Unmarshal([]byte(*key.SecretString), &secretData)
	if err != nil {
		fmt.Println("Error al deserializar el secreto:", err.Error())
		return secretData, err
	}

	fmt.Println("Secreto deserializado:", secretData)
	return secretData, nil
}