package repositories

import (
	"fmt"
	"go-react/server/models"
)

type StudentRepository interface {
	FindAllStudent() *[]models.Student
	FindStudentByID(stdID uint) *models.Student
	FindStudentByNIM(nim string) *models.Student
	CreateStudent(std *models.Student) *models.Student
	UpdateStudent(std *models.Student) *models.Student
	DeleteStudent(stdID uint) *models.Student
}

func (r *repository) FindAllStudent() *[]models.Student {
	return &models.StudentData
}

func (r *repository) FindStudentByID(stdID uint) *models.Student {
	var result models.Student

	// find student by id
	for _, std := range models.StudentData {
		if std.ID == stdID {
			result = std
		}
	}

	return &result
}

func (r *repository) FindStudentByNIM(nim string) *models.Student {
	var result models.Student

	// find student by nim
	for _, std := range models.StudentData {
		if std.NIM == nim {
			result = std
		}
	}

	return &result
}

func (r *repository) CreateStudent(std *models.Student) *models.Student {
	models.StudentData = append(models.StudentData, *std)

	printLog(*std, "created")

	return std
}

func (r *repository) UpdateStudent(std *models.Student) *models.Student {
	for i, data := range models.StudentData {
		if data.ID == std.ID {
			models.StudentData[i] = *std
		}
	}

	printLog(*std, "updated")

	return std
}

func (r *repository) DeleteStudent(stdID uint) *models.Student {
	var (
		newStudentData []models.Student
		deletedStudent models.Student
	)

	for _, std := range models.StudentData {
		if std.ID != stdID {
			newStudentData = append(newStudentData, std)
		} else if std.ID == stdID {
			deletedStudent = std
		}
	}

	models.StudentData = newStudentData

	printLog(deletedStudent, "deleted")

	return &deletedStudent
}

func printLog(std models.Student, action string) {
	fmt.Println("Student with nim -> " + std.NIM + " has been " + action + " successfully !")
}
