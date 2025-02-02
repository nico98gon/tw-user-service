package utils

// joinErrors convierte un slice de errores en un solo string con saltos de línea
func JoinErrors(errors []string) string {
	result := ""
	for _, err := range errors {
		result += "- " + err + "\n"
	}
	return result
}