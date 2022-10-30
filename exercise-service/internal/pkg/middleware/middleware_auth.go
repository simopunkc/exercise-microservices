package middleware

import (
	"exercise-service/internal/app/container"
	"exercise-service/internal/app/domain"
	"exercise-service/internal/pkg/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationJWT(c *fiber.Ctx) error {
	authHeader := string(c.Request().Header.Peek("Authorization"))
	if authHeader == "" {
		return c.Status(400).JSON(domain.MiddlewareError{
			Hash:  "GMkeVzQiMWPH",
			Error: "unauthorize",
		})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(400).JSON(domain.MiddlewareError{
			Hash:  "GMxxg9BZl0qo",
			Error: "unauthorize",
		})
	}

	auths := strings.Split(authHeader, " ")
	utilJwt := util.NewUtilJwt()
	data, err := utilJwt.DecriptJWT(auths[1])
	if err != nil {
		return c.Status(400).JSON(domain.MiddlewareError{
			Hash:  "GMg0ijXJCAZ2",
			Error: "unauthorize",
		})
	}
	userIDinterface, ok := data["user_id"]
	if !ok {
		return c.Status(400).JSON(domain.MiddlewareError{
			Hash:  "GMNcUFjL7ODe",
			Error: "invalid user id",
		})
	}
	userID := int64(userIDinterface.(float64))
	userService := container.NewContainerUser()
	if !userService.IsUserExists(c.Context(), userID) {
		return c.Status(401).JSON(domain.MiddlewareError{
			Hash:  "GMWt2r2vJj6M",
			Error: "user not exists",
		})
	}
	c.Locals("user_id", userID)
	return c.Next()
}
