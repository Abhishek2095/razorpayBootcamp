package Routes

import (
	"q2/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetStudents)
		grp1.POST("user", Controllers.CreateStudent)
		grp1.GET("user/:id", Controllers.GetStudentByID)
		grp1.PUT("user/:id", Controllers.UpdateStudent)
		grp1.DELETE("user/:id", Controllers.DeleteStudent)
	}

	grp2 := r.Group("/marks-api")
	{
		grp2.GET("marks", Controllers.GetMarks)
		grp2.POST("marks", Controllers.CreateMarks)
	}

	return r
}

// {
//     "id": 1,
//     "firstname": "abcd",
//     "lastname": "efgh",
//     "address": "Street 5",
//     "dateofbirth": "1987-10-12"
// }

// {
//     "studentid": 6,
//     "subject": "English",
//     "marks": 85
// }
