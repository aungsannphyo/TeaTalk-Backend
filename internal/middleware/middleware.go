package middleware

import (
	"net/http"
	"strings"

	"github.com/aungsannphyo/ywartalk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")

	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": " Authorization header missing or malformed"})
		return
	}

	token := strings.TrimPrefix(header, "Bearer ")
	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
