package users

import (
	"errors"
	"user-service/internal/utils"
)

// LoginValidations verifica las credenciales de inicio de sesión
func LoginValidations(u User) error {
	var errorMessages []string

	email := u.Email
	password := u.Password

	if err := validateEmail(email); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validatePassword(password); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if len(errorMessages) > 0 {
		return errors.New("Errores de validación en el login:\n" + utils.JoinErrors(errorMessages))
	}

	return nil
}
