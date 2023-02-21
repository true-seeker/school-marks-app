package controllers

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	db "school-marks-app/api/db/models"
	"school-marks-app/api/db/models/validators"
	"strconv"
)

type StudentController struct{}

func (s StudentController) GetById(c *gin.Context) {
	var studentModel db.Student
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	teacher, webErr := studentModel.GetById(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, teacher)
	return
}

func (s StudentController) Create(c *gin.Context) {
	var student db.Student
	if validationErr := c.BindJSON(&student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}
	if validationErr := validators.ValidateStudentCreate(student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := student.Create()
	if webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (s StudentController) Update(c *gin.Context) {
	var student db.Student
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if validationErr := c.BindJSON(&student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	student.ID = id

	if validationErr := validators.ValidateStudentUpdate(student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := student.Update()
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (s StudentController) Delete(c *gin.Context) {
	var student db.Student
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if webErr := student.Delete(id); webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (s StudentController) BulkCreate(c *gin.Context) {
	var studentModel db.Student
	var csvStudents []db.Student
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

		csvStudents = append(csvStudents, db.Student{
			Name:       line[0],
			Surname:    line[1],
			Patronymic: line[2],
			ClassID:    uint(classId),
		})
	}
	newIds, webErr := studentModel.BulkCreate(csvStudents)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, newIds)

}
