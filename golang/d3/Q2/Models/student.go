package Models

import (
	"fmt"
	"q2/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllStudents Fetch all user data
func GetAllStudents(student *[]Student) (err error) {

	// rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
	// rows, err := Config.DB.Table("students").Select("students.id, students.firstname, marks.subject, marks.marks").Joins("inner join marks on students.id = marks.studentid").Rows()
	// if err = Config.DB.Table("students").Select("students.id, marks.subject, marks.marks").Joins("inner join marks on students.id = marks.student_id").Find(student).Error; err != nil {
	if err = Config.DB.Find(student).Error; err != nil {
		return err
	}
	return nil
}

//CreateStudent ... Insert New data
func CreateStudent(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

//UpdateStudent
func UpdateUser(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

//DeleteStudent
func DeleteUser(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}
