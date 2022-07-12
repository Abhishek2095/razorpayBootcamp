package Validate

import (
	"errors"
	"regexp"
	"studentAPI/Models"
)

func ValidateStudent(student *Models.Student) error {
	if err := validateName(student.First_name); err != nil {
		return err
	}
	if err := validateName(student.Last_name); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	matched, err := regexp.MatchString("^[A-Z][a-z]+$", name)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("invalid name")
	}
	return nil
}
