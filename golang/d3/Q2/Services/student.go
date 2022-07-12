package Services

import (
	"q2/Models"
	"q2/Services/Parse"
	"q2/Services/Validate"
)

func CreateStudent(requestJSON map[string]string) error {
	var student Models.Student

	if studentId, err := Parse.ParseId(requestJSON); err != nil {
		return err
	} else {

		student.Id = studentId
	}

	if firstName, lastName, err := Parse.ParseName(requestJSON); err != nil {
		return err
	} else {
		if student.FirstName, err = Validate.ValidateName(firstName); err != nil {
			return err
		}
		if student.LastName, err = Validate.ValidateName(lastName); err != nil {
			return err
		}
	}

	if address, err := Parse.ParseAddress(requestJSON); err != nil {
		return err
	} else {
		student.Address = address
	}

	if dateOfBirth, err := Parse.ParseDateOfBirth(requestJSON); err != nil {
		return err
	} else {
		if err = Validate.ValidateDateOfBirth(dateOfBirth); err != nil {
			return err
		} else {
			student.DateOfBirth = dateOfBirth
		}
	}

	student.MarksList = nil

	if err := Models.CreateStudent(&student); err != nil {
		return err
	}
	return nil
}
