package middleware

import (
	"context"
	"exercise-service/internal/user"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithJWT(us *user.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"message": "unauthorize",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, gin.H{
				"message": "unauthorize",
			})
			c.Abort()
			return
		}

		auths := strings.Split(authHeader, " ")
		data, err := us.DecriptJWT(auths[1])
		if err != nil {
			c.JSON(401, gin.H{
				"message": "unauthorize",
			})
			c.Abort()
			return
		}
		userIDinterface, ok := data["user_id"]
		if !ok {
			c.JSON(401, gin.H{
				"message": "invalid data",
			})
			c.Abort()
			return
		}
		userID := int(userIDinterface.(float64))
		if !us.IsUserExists(c.Request.Context(), userID) {
			c.JSON(401, gin.H{
				"message": "user not exists",
			})
			c.Abort()
			return
		}
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", userID)
		c.Request = c.Request.WithContext(ctxUserID)
		c.Next()
	}
}
