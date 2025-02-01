package db

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pass string) (string, error) {
	costStr := os.Getenv("BCRYPT_COST")
	if costStr == "" {
		fmt.Println("BCRYPT_COST env no configurada")
		return "", nil
	}

	cost, err := strconv.Atoi(costStr)
	if err != nil {
		fmt.Println("Error al convertir BCRYPT_COST a número:", err)
		return "", err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		return err.Error(), err
	}

	return string(bytes), nil
}