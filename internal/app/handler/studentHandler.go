package handler

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/service"
	"school-marks-app/internal/app/validator"
	"strconv"
)

type StudentHandler struct {
	service service.Student
}

func NewStudentHandler(service service.Student) *StudentHandler {
	return &StudentHandler{service: service}
}

func (s *StudentHandler) GetById(c *gin.Context) {
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	teacher, webErr := s.service.GetById(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, teacher)
	return
}

func (s *StudentHandler) Create(c *gin.Context) {
	student := &entity.Student{}
	if validationErr := c.BindJSON(&student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}
	if validationErr := validator.ValidateStudentCreate(student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newStudent, webErr := s.service.Create(student)
	if webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newStudent)
}

func (s *StudentHandler) Update(c *gin.Context) {
	student := &entity.Student{}
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if validationErr := c.BindJSON(&student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	student.ID = id

	if validationErr := validator.ValidateStudentUpdate(student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newStudent, webErr := s.service.Update(student)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newStudent)
}

func (s *StudentHandler) Delete(c *gin.Context) {
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if webErr := s.service.Delete(id); webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (s *StudentHandler) BulkCreate(c *gin.Context) {
	var csvStudents []entity.Student
	filePtr, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	file, err := filePtr.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if len(records) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Empty csv")})
		return
	}

	if len(records[0]) != 4 || records[0][0] != "name" || records[0][1] != "surname" ||
		records[0][2] != "patronymic" || records[0][3] != "class_id" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Wrong csv format")})
		return
	}

	for rowId, line := range records[1:] {
		classId, err := strconv.Atoi(line[3])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Field class_id must be integer on line %d", rowId+1)})
			return
		}
		csvStudent := entity.Student{
			Name:       line[0],
			Surname:    line[1],
			Patronymic: line[2],
			ClassID:    uint(classId),
		}
		if validationErr := validator.ValidateStudentCreate(&csvStudent); validationErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error() + fmt.Sprintf(" on line %d", rowId+1)})
			return
		}
		csvStudents = append(csvStudents, csvStudent)
	}
	newIds, webErr := s.service.BulkCreate(csvStudents)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, newIds)
}
