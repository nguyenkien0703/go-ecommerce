package auth

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func ExtractBearerToken(c *gin.Context) (string, bool) {

	// Authorization: Bearer token
	// token
	log.Println("c-----", c.Request)
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer") {
		return strings.TrimPrefix(authHeader, "Bearer "), true
	}
	return "", false
}
