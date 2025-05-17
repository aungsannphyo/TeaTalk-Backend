package middleware

import (
	"net/http"
	"strings"

	"github.com/aungsannphyo/ywartalk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {

	header := c.Request.Header.Get("Authorization")
	var token string

	if header != "" && strings.HasPrefix(header, "Bearer ") {
		token = strings.TrimPrefix(header, "Bearer ")
	} else {
		token = c.Query("token")
	}
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token missing in header or query"})
		return
	}

	userID, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	c.Set("userID", userID)
	c.Next()
}
