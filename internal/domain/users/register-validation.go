package users

import (
	"errors"
	"user-service/internal/utils"
)

// Validations verifica todas las validaciones de un usuario y devuelve un error consolidado si hay fallos
func RegisterValidations(u User) error {
	var errorMessages []string

	if err := validateName(u.Name, true); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validateLastName(u.LastName, true); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validateEmail(u.Email); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validateBirthdate(u.Birthdate); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validatePassword(u.Password); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validateURL(u.Avatar, true); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validateURL(u.Banner, true); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validateURL(u.WebSite, true); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validateLocation(u.Location, true); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if len(errorMessages) > 0 {
		return errors.New("Errores de validación:\n" + utils.JoinErrors(errorMessages))
	}

	return nil
}
