package success

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OkResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func CreateResponse(c *gin.Context, message string) {
	c.JSON(http.StatusCreated, gin.H{"message": message})
}
