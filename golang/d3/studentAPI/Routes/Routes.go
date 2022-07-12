package Routes

import (
	"studentAPI/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/student-api")
	{
		grp1.GET("student/", Controllers.GetStudents)
		grp1.POST("student/:id/", Controllers.CreateStudent)
		grp1.GET("student/:id/", Controllers.GetStudentByID)
		grp1.PUT("student/:id/", Controllers.UpdateInfo)
		grp1.DELETE("student/:id/", Controllers.DeleteStudent)

		grp1.GET("marks/", Controllers.GetAllMarks)
		grp1.GET("marks/:student_id/", Controllers.GetMarksByID)
		grp1.PUT("marks/:student_id/:id", Controllers.UpdateMarks)
		grp1.DELETE("marks/:student_id/:id", Controllers.DeleteMarks)
		grp1.DELETE("marks/:student_id/", Controllers.DeleteAllMarks)
	}
	return r
}
