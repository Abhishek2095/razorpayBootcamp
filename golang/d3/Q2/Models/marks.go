package Models

import (
	"q2/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllMarks(marks *[]Marks) (err error) {
	if err = Config.DB.Find(marks).Error; err != nil {
		return err
	}
	return nil
}

func CreateMarks(marks *Marks) (err error) {
	if err = Config.DB.Create(marks).Error; err != nil {
		return err
	}
	return nil
}
