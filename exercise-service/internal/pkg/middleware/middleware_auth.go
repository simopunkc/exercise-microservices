package middleware

import (
	"errors"
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
			Hash:  "",
			Error: errors.New("unauthorize"),
		})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(400).JSON(domain.MiddlewareError{
			Hash:  "",
			Error: errors.New("unauthorize"),
		})
	}

	auths := strings.Split(authHeader, " ")
	data, err := util.DecriptJWT(auths[1])
	if err != nil {
		return c.Status(400).JSON(domain.MiddlewareError{
			Hash:  "",
			Error: errors.New("unauthorize"),
		})
	}
	userIDinterface, ok := data["user_id"]
	if !ok {
		return c.Status(400).JSON(domain.MiddlewareError{
			Hash:  "",
			Error: errors.New("invalid user id"),
		})
	}
	userID := int64(userIDinterface.(float64))
	userService := container.NewContainerUser()
	if !userService.IsUserExists(c.Context(), userID) {
		return c.Status(401).JSON(domain.MiddlewareError{
			Hash:  "",
			Error: errors.New("user not exists"),
		})
	}
	c.Locals("user_id", userID)
	return c.Next()
}
