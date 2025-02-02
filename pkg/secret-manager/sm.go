package secretmanager

import (
	"encoding/json"
	"fmt"
	domain "user-service/internal/domain/sm"
	awsSession "user-service/pkg/aws"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// GetSecret obtiene las claves secretas de AWS
func GetSecret(secretName string) (domain.Secret, error) {
	var secretData domain.Secret
	fmt.Println("> Se pide secreto:", secretName)

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