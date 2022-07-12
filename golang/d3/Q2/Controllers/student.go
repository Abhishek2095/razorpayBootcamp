package Controllers

import (
	"fmt"
	"net/http"
	"q2/Models"

	"github.com/gin-gonic/gin"
)

// CreateStudent
func CreateStudent(c *gin.Context) {
	var student Models.Student
	// jsonData, _ := c.GetRawData()
	// fmt.Println("\n---------\n", string(jsonData), "\n----------")
	c.BindJSON(&student)
	err := Models.CreateStudent(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// func CreateStudentHandler(c *gin.Context) {
// 	requestJSON := make(map[string]string)

// 	byteData, err := c.GetRawData()
// 	if err != nil {
// 		fmt.Println("Error while reading data from request : ", err)
// 	}

// 	if err = json.Unmarshal(byteData, &requestJSON); err != nil {
// 		fmt.Println("Error while Unmarshaling : ", err)
// 	}

// 	fmt.Println("---------")
// 	fmt.Println(requestJSON)
// 	fmt.Println("---------")
// 	Services.CreateStudent(requestJSON)
// }

//GetStudent
func GetStudents(c *gin.Context) {
	var students []Models.Student
	err := Models.GetAllStudents(&students)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

//GetStudentByID
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetUserByID(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//UpdateStudent
func UpdateStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models.UpdateUser(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//DeleteStudent
func DeleteStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
