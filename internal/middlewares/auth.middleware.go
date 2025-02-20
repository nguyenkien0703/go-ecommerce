package middlewares

import (
	"context"
	"example.com/go-ecommerce-backend-api/internal/utils/auth"
	"github.com/gin-gonic/gin"
	"log"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the request url path
		uri := c.Request.URL.Path
		log.Println("uri request---", uri)

		// check headers authorization
		jwtToken, valid := auth.ExtractBearerToken(c)
		if !valid {
			c.AbortWithStatusJSON(401, gin.H{
				"code":        401,
				"err":         "Unauthorized",
				"description": "",
			})
			return
		}

		//validate jwt token by subject
		claims, err := auth.VerifyTokenSubject(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"code":        401,
				"err":         "Invalid token",
				"description": "",
			})
			return
		}

		// update claims to context
		log.Println("claims:::UUID", claims.Subject) // 11clitoken....
		ctx := context.WithValue(c.Request.Context(), "subjectUUID", claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}

}
