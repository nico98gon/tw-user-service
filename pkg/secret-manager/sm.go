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
	fmt.Println("> Se pide secreto"+ secretName)

	svc := secretsmanager.NewFromConfig(awsSession.Cfg)
	key, err := svc.GetSecretValue(awsSession.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secretData)
	fmt.Println(" > Obtención secreto OK"+ secretName)

	return secretData, nil
}