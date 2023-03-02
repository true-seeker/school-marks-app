package handler

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"school-marks-app/internal/app/model/entity"
	db "school-marks-app/internal/app/model/input"
	"school-marks-app/internal/app/service"
	"school-marks-app/internal/app/validator"
	"strconv"
)

type ClassHandler struct {
	service             service.Class
	academicYearService service.AcademicYear
	schoolClassService  service.SchoolClass
}

func NewClassHandler(service service.Class, academicYearService service.AcademicYear, schoolClassService service.SchoolClass) *ClassHandler {
	return &ClassHandler{service: service, academicYearService: academicYearService, schoolClassService: schoolClassService}
}

func (cl *ClassHandler) GetById(c *gin.Context) {
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	class, webErr := cl.service.GetById(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	if class.Id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("class with id %d does not exist", id))
		return
	}

	c.JSON(http.StatusOK, class)
	return
}

func (cl *ClassHandler) Create(c *gin.Context) {
	var newClass *entity.Class

	if validationErr := c.BindJSON(&newClass); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}
	if validationErr := validator.ValidateClassCreate(newClass); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	class, webErr := cl.service.Create(newClass)
	if webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, class)
}

func (cl *ClassHandler) Update(c *gin.Context) {
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}
	class := &entity.Class{}
	if validationErr := c.BindJSON(&class); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}
	class.ID = id

	if validationErr := validator.ValidateClassUpdate(class); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newClass, webErr := cl.service.Update(class)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newClass)
}

func (cl *ClassHandler) Delete(c *gin.Context) {
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	webErr := cl.service.Delete(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (cl *ClassHandler) BulkCreate(c *gin.Context) {
	csvClasses := make([]entity.Class, 0)
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

	if len(records[0]) != 4 || records[0][0] != "teacher_id" || records[0][1] != "year_id" ||
		records[0][2] != "school_class_id" || records[0][3] != "letter" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Wrong csv format")})
		return
	}

	for rowId, line := range records[1:] {
		teacherId, err := strconv.Atoi(line[0])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Field teacher_id must be integer on line %d", rowId+1)})
			return
		}
		yearId, err := strconv.Atoi(line[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Field year_id must be integer on line %d", rowId+1)})
			return
		}
		schoolClassId, err := strconv.Atoi(line[2])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Field school_class_id must be integer on line %d", rowId+1)})
			return
		}

		csvClass := &entity.Class{
			TeacherID:     uint(teacherId),
			YearID:        uint(yearId),
			SchoolClassId: uint(schoolClassId),
			Letter:        line[3],
		}
		if validationErr := validator.ValidateClassCreate(csvClass); validationErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error() + fmt.Sprintf(" on line %d", rowId+1)})
			return
		}
		csvClasses = append(csvClasses, *csvClass)
	}
	newIds, webErr := cl.service.BulkCreate(&csvClasses)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, newIds)
}

func (cl *ClassHandler) Transfer(c *gin.Context) {
	var classTransferInput db.ClassTransferInput

	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if validationErr := c.BindJSON(&classTransferInput); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	if _, webErr := cl.academicYearService.GetById(classTransferInput.NewAcademicYearId); webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, webErr.Err.Error())
		return
	}
	if _, webErr := cl.schoolClassService.GetById(classTransferInput.NewSchoolClassId); webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, webErr.Err.Error())
		return
	}

	if webErr := cl.service.Transfer(classTransferInput, id); webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
}
