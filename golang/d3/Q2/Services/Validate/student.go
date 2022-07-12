package Validate

import (
	"regexp"
)

func ValidateName(name string) (string, error) {
	_, err := regexp.MatchString("^[A-Z][a-z]+$", name)
	if err != nil {
		return "", err
	}

	return name, nil
}
