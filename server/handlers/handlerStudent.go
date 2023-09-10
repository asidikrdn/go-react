package handlers

import (
	"go-react/server/dto"
	"go-react/server/models"
	"go-react/server/pkg/helpers"
	"go-react/server/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerStudent struct {
	StudentRepository repositories.StudentRepository
}

func HandlerStudent(stdRepository repositories.StudentRepository) *handlerStudent {
	return &handlerStudent{stdRepository}
}

func convertStudentResponse(std *models.Student) *dto.StudentResponse {
	return &dto.StudentResponse{
		ID:       std.ID,
		NIM:      std.NIM,
		Fullname: std.Fullname,
		Majority: std.Majority,
		Address:  std.Address,
		Image:    std.Pict,
	}
}

func convertMultipleStudentResponse(std *[]models.Student) *[]dto.StudentResponse {
	var students []dto.StudentResponse

	for _, m := range *std {
		students = append(students, *convertStudentResponse(&m))
	}

	return &students
}

func (h *handlerStudent) CreateStudent(c *gin.Context) {
	var request dto.StudentRequest

	// get request data
	err := c.ShouldBind(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// check is student already registered
	std := h.StudentRepository.FindStudentByNIM(request.NIM)
	if std.ID != 0 {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Student already registered",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create new student
	student := models.Student{
		ID:       helpers.GetLastID() + 1,
		NIM:      request.NIM,
		Fullname: request.Fullname,
		Majority: request.Majority,
		Address:  request.Address,
	}

	// get image from context
	image, ok := c.Get("image")
	if ok {
		student.Pict = image.(string)
	}

	// save new student data to database
	addedStudent := h.StudentRepository.CreateStudent(&student)

	// reload data
	student = *h.StudentRepository.FindStudentByID(addedStudent.ID)

	// send response
	response := dto.Result{
		Status:  http.StatusCreated,
		Message: "OK",
		Data:    convertStudentResponse(&student),
	}
	c.JSON(http.StatusCreated, response)
}

func (h *handlerStudent) FindAllStudent(c *gin.Context) {
	// get student data
	students := h.StudentRepository.FindAllStudent()

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertMultipleStudentResponse(students),
	}
	c.JSON(http.StatusOK, response)
}

func (h *handlerStudent) FindStudentByID(c *gin.Context) {
	// get student id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get student data
	student := h.StudentRepository.FindStudentByID(uint(id))
	if student == nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: notFound,
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertStudentResponse(student),
	}
	c.JSON(http.StatusOK, response)
}

func (h *handlerStudent) UpdateStudentByID(c *gin.Context) {
	var request dto.StudentRequest

	// get request data
	err := c.ShouldBind(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get student id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get student data
	student := h.StudentRepository.FindStudentByID(uint(id))
	if student == nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: notFound,
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	if student.NIM != request.NIM {
		student.NIM = request.NIM
	}

	if student.Fullname != request.Fullname {
		student.Fullname = request.Fullname
	}

	if student.Majority != request.Majority {
		student.Majority = request.Majority
	}

	if student.Address != request.Address {
		student.Address = request.Address
	}

	// update image
	image, ok := c.Get("image")
	if ok {
		if student.Pict != "" {
			if !helpers.DeleteFile(student.Pict) {
				log.Println(err.Error())
			}
		}

		student.Pict = image.(string)
	}

	// update student data
	student = h.StudentRepository.UpdateStudent(student)

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertStudentResponse(student),
	}
	c.JSON(http.StatusOK, response)
}

func (h *handlerStudent) DeleteStudentByID(c *gin.Context) {
	// get student id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// delete student data
	student := h.StudentRepository.DeleteStudent((uint(id)))
	if student == nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: notFound,
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertStudentResponse(student),
	}
	c.JSON(http.StatusOK, response)
}

var notFound = "Mahasiwa not found"
