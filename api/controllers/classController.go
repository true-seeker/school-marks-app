package controllers

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	db "school-marks-app/api/db/dto"
	db2 "school-marks-app/api/db/models"
	"school-marks-app/api/db/models/validators"
	"strconv"
)

type ClassController struct{}

func (cl ClassController) GetById(c *gin.Context) {
	var classControllerModel db2.Class
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	class, webErr := classControllerModel.GetById(id)

	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
	return
}

func (cl ClassController) Create(c *gin.Context) {
	var class db2.Class

	if validationErr := c.BindJSON(&class); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}
	if validationErr := validators.ValidateClassCreate(class); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := class.Create()
	if webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (cl ClassController) Update(c *gin.Context) {
	var class db2.Class

	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if validationErr := c.BindJSON(&class); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	class.ID = id

	if validationErr := validators.ValidateClassUpdate(class); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := class.Update()
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (cl ClassController) Delete(c *gin.Context) {
	var class db2.Class
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	webErr := class.Delete(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (cl ClassController) BulkCreate(c *gin.Context) {
	var classModel db2.Class
	var csvClasses []db2.Class
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

	if len(records[0]) != 3 || records[0][0] != "teacher_id" || records[0][1] != "year_id" ||
		records[0][2] != "school_class_id" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Wrong csv format")})
		return
	}

	for rowId, line := range records[1:] {
		teacherId, err := strconv.Atoi(line[0])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Field teacher_id must be integer on line %d", rowId+1)})
			return
		}
		yaerId, err := strconv.Atoi(line[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Field year_id must be integer on line %d", rowId+1)})
			return
		}
		schoolClassId, err := strconv.Atoi(line[2])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Field school_class_id must be integer on line %d", rowId+1)})
			return
		}

		csvClass := db2.Class{
			TeacherID:     uint(teacherId),
			YearID:        uint(yaerId),
			SchoolClassId: uint(schoolClassId),
		}
		if validationErr := validators.ValidateClassCreate(csvClass); validationErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error() + fmt.Sprintf(" on line %d", rowId+1)})
			return
		}
		csvClasses = append(csvClasses, csvClass)
	}
	newIds, webErr := classModel.BulkCreate(csvClasses)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, newIds)
}

func (cl ClassController) Transfer(c *gin.Context) {
	var classTransferInput db.ClassTransferInput

	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}
	class := db2.Class{
		Model: gorm.Model{ID: id},
	}

	if validationErr := c.BindJSON(&classTransferInput); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	if webErr := class.Transfer(classTransferInput); webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
}
