package users

import (
	"errors"
	"regexp"
	"time"
	"unicode"
	"user-service/internal/utils"
)

func validateName(name string) error {
	if len(name) == 0 {
		return errors.New("el nombre es requerido")
	}
	if !regexp.MustCompile(`^[a-zA-ZáéíóúÁÉÍÓÚñÑ\s]+$`).MatchString(name) {
		return errors.New("el nombre solo puede contener letras y espacios")
	}
	return nil
}

func validateLastName(lastName string) error {
	if len(lastName) == 0 {
		return errors.New("el apellido es requerido")
	}
	if !regexp.MustCompile(`^[a-zA-ZáéíóúÁÉÍÓÚñÑ\s]+$`).MatchString(lastName) {
		return errors.New("el apellido solo puede contener letras y espacios")
	}
	return nil
}

func validateEmail(email string) error {
	if len(email) == 0 {
		return errors.New("el email es requerido")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("el email no tiene un formato válido")
	}
	return nil
}

func validateBirthdate(birthdate time.Time) error {
	now := time.Now()
	if birthdate.After(now) {
		return errors.New("la fecha de nacimiento no puede ser en el futuro")
	}
	minimumAge := now.AddDate(-13, 0, 0)
	if birthdate.After(minimumAge) {
		return errors.New("debes tener al menos 13 años")
	}
	return nil
}

func validatePassword(password string) error {
	var errorMessages []string

	if len(password) < 6 {
		errorMessages = append(errorMessages, "la contraseña debe tener al menos 6 caracteres")
	}

	var hasUpper, hasLower, hasDigit bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if !hasLower {
		errorMessages = append(errorMessages, "la contraseña debe contener al menos una minúscula")
	}
	if !hasUpper {
		errorMessages = append(errorMessages, "la contraseña debe contener al menos una mayúscula")
	}
	if !hasDigit {
		errorMessages = append(errorMessages, "la contraseña debe contener al menos un número")
	}

	if len(errorMessages) > 0 {
		return errors.New(utils.JoinErrors(errorMessages))
	}

	return nil
}

func validateURL(url string) error {
	if len(url) == 0 {
		return nil
	}
	urlRegex := regexp.MustCompile(`^https?:\/\/[^\s]+$`)
	if !urlRegex.MatchString(url) {
		return errors.New("la URL no es válida")
	}
	return nil
}

func validateLocation(location string) error {
	if len(location) > 100 {
		return errors.New("la ubicación no puede tener más de 100 caracteres")
	}
	return nil
}