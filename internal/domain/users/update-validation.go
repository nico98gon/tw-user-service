package users

import (
	"errors"
	"user-service/internal/utils"
)

func UpdateValidations(u User) error {
	var errorMessages []string

	if err := validateName(u.Name, false); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validateLastName(u.LastName, false); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validateURL(u.WebSite); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if err := validateLocation(u.Location); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if len(errorMessages) > 0 {
		return errors.New("Errores de validación:\n" + utils.JoinErrors(errorMessages))
	}

	return nil
}