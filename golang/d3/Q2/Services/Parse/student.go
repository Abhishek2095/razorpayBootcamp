package Parse

import (
	"errors"
	"strconv"
	"time"
)

func ParseId(requestJSON map[string]string) (uint, error) {
	var studentId uint = 0

	studentIdString, found := requestJSON["id"]
	if !found {
		return studentId, errors.New("missing Key")
	}

	studentIdInt, err := strconv.Atoi(studentIdString)
	if err != nil {
		return studentId, err
	} else if studentIdInt < 1 {
		return studentId, errors.New("id cannot be negative")
	}

	studentId = uint(studentIdInt)
	return studentId, nil
}

func ParseName(requestJSON map[string]string) (string, string, error) {
	var firstName string = ""
	var lastName string = ""
	var found bool

	firstName, found = requestJSON["firstname"]
	if !found {
		return "", "", errors.New("first name not found")
	}
	lastName, found = requestJSON["lastname"]
	if !found {
		return "", "", errors.New("last name not found")
	}

	return firstName, lastName, nil
}

func ParseAddress(requestJSON map[string]string) (string, error) {
	address, found := requestJSON["address"]
	if !found {
		return "", errors.New("address not found")
	}

	return address, nil
}

func ParseDateOfBirth(requestJSON map[string]string) (time.Time, error) {
	dateOfBirthString, found := requestJSON["dateofbirth"]
	if !found {
		return time.Now(), errors.New("date of birth not found")
	}

	if dateOfBirth, err := time.Parse("2006-01-02", dateOfBirthString); err != nil {
		return time.Now(), err
	} else {
		return dateOfBirth, nil
	}
}
