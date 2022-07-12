package Services

import (
	"errors"
	"regexp"
	"strconv"
	"studentAPI/Models"
)

var (
	ErrorInvalidId   = errors.New("id cannot be negative")
	ErrorInvalidName = errors.New("invalid name")
)

func ValidateStudent(student *Models.Student) error {
	if err := validateID(student.Id); err != nil {
		return err
	}
	if err := validateName(student.First_name); err != nil {
		return err
	}
	if err := validateName(student.Last_name); err != nil {
		return err
	}
	return nil

}

func validateID(s string) error {
	if id, err := strconv.Atoi(s); err != nil {
		return err
	} else if id <= 0 {
		return ErrorInvalidId
	}

	return nil
}

func validateName(name string) error {
	var IsOnlyLetters = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	if !IsOnlyLetters(name) {
		return ErrorInvalidName
	}
	return nil
}
