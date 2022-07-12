package Controllers

import (
	"fmt"
	"net/http"
	"q2/Models"

	"github.com/gin-gonic/gin"
)

//GetStudents
func GetMarks(c *gin.Context) {
	var marks []Models.Marks
	err := Models.GetAllMarks(&marks)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}

//CreateStudent
func CreateMarks(c *gin.Context) {
	var marks Models.Marks
	c.BindJSON(&marks)
	err := Models.CreateMarks(&marks)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}
