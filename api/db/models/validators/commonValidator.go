package validators

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ValidateAndReturnId(c *gin.Context, id string) (uint, error) {
	if c.Param("id") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing field \"id\""})
		return 0, errors.New("")
	}

	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Field id must be integer"})
		return 0, errors.New("")
	}
	return uint(idInt), nil
}
