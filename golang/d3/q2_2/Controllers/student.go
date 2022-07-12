package Controllers

import (
	"fmt"
	"net/http"
	"studentAPI/Models"
	"studentAPI/Services"

	"github.com/gin-gonic/gin"
)

//GetStudents ... Get all students
func GetStudents(c *gin.Context) {
	var students []Models.Student
	err := Models.GetAllStudents(&students)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

//CreateStudent ... Create User
func CreateStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	if err := Services.ValidateStudent(&student); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		err := Models.CreateStudent(&student)

		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, student)
		}
	}
}

//GetStudentByID ... Get the user by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetStudentByID(&student, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//UpdateInfo ... Update the user information
func UpdateInfo(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentByID(&student, id)

	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models.UpdateInfo(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//DeleteStudent ... Delete the user
func DeleteStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&student, id)

	// ********   ALSO DELETE THE MARKS TABLE ASSOCIATED WITH THE STUDENT   ********

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var marks []Models.Marks
		err = Models.DeleteAllMarks(&marks, id)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"All marks for Student " + id: " are deleted"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}

}

//GetAllMarks ... Get all marks
func GetAllMarks(c *gin.Context) {
	var marks []Models.Marks // Check for availability of marks
	err := Models.GetAllMarks(&marks)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}

//GetMarksByID ... Get the user by id
func GetMarksByID(c *gin.Context) {

	studentId := c.Params.ByName("student_id")
	var student Models.Student // Check for availability of student
	err := Models.GetStudentByID(&student, studentId)
	if err == nil {
		var marks Models.Marks // Check for availability of marks
		id := c.Params.ByName("id")
		err = Models.GetMarksByID(&marks, id)

		if err == nil {
			var empty = Models.Marks{}
			if marks != empty {
				c.JSON(http.StatusOK, marks)
			} else {
				c.JSON(http.StatusOK, "No Marks Data found for the student")
			}
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	} else {
		c.JSON(http.StatusOK, "Student ID Not Found")
	}

}

//UpdateMarks ... Update the user information
func UpdateMarks(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")

	studentId := c.Params.ByName("student_id")
	err := Models.GetStudentByID(&student, studentId) // Check for availability of student
	if err == nil {
		var marks Models.Marks // Check for availability of marks
		c.BindJSON(&marks)

		err = Models.UpdateMarks(&marks, id)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, marks)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}

}

//DeleteMarks ... Delete the marks by id		*********** 1 More case-----No result found for marks for the id
func DeleteMarks(c *gin.Context) {
	studentId := c.Params.ByName("student_id")
	id := c.Params.ByName("id")
	var student Models.Student // Check for availability of student
	err := Models.GetStudentByID(&student, studentId)
	if err == nil {
		var marks Models.Marks // Check for availability of marks
		err = Models.GetMarksByID(&marks, studentId)
		if err == nil {
			err = Models.DeleteMarks(&marks, id)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"Marks for Student " + studentId: "and Marks ID " + id + " are deleted"})
			} else {
				c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			}
		} else {
			c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		}

	} else {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}

}

func DeleteAllMarks(c *gin.Context) {
	studentId := c.Params.ByName("student_id")
	var student Models.Student // Check for availability of student
	err := Models.GetStudentByID(&student, studentId)
	if err == nil {
		var marks []Models.Marks // Check for availability of marks
		err = Models.DeleteAllMarks(&marks, studentId)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"All marks for Student " + studentId: " are deleted"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}

}
