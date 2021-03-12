package common 

import (
	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Keys["currentUser"].(*UserPayload); !ok {
			c.Error(NewNotAuthorizedError())
			c.Abort()
		}

		c.Next()
	}
}
