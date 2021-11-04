package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hjunor/api-rest-golang.git/services"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_shema = "Bearer "
		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatus(401)
			return
		}

		token := header[len(Bearer_shema):]

		if !services.NewJwtServices().ValidateToken(token) {
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
