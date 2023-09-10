package routes

import (
	"go-react/server/handlers"
	"go-react/server/repositories"

	"github.com/gin-gonic/gin"
)

func Student(r *gin.RouterGroup) {
	stdRepository := repositories.MakeRepository()
	h := handlers.HandlerStudent(stdRepository)

	studentByID := "/students/:id"

	r.POST("/students", h.CreateStudent)
	r.GET("/students", h.FindAllStudent)
	r.GET(studentByID, h.FindStudentByID)
	r.PATCH(studentByID, h.UpdateStudentByID)
	r.DELETE(studentByID, h.DeleteStudentByID)
}
