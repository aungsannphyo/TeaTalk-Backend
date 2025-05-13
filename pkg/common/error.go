package common

import (
	"net/http"

	v "github.com/aungsannphyo/ywartalk/pkg/validator"
	"github.com/gin-gonic/gin"
)

func BadRequestResponse(c *gin.Context, err error) {
	if ve, ok := err.(v.ValidationErrors); ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": ve,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func InternalServerResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func UnauthorizedResponse(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
}

func NotFoundResponse(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
}

func ConfictResponse(c *gin.Context, err error) {
	c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
}
