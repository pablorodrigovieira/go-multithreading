package utils

import (
	"regexp"
	"strings"
)

func ValidateCEP(cep string) (string, bool) {
	// Remove any dash or spaces
	clean := strings.ReplaceAll(cep, "-", "")
	clean = strings.TrimSpace(clean)

	// Check if exactly 8 digits
	match, _ := regexp.MatchString(`^\d{8}$`, clean)
	return clean, match
}
