package route

import (
	"os"
	"user-service/internal/app/container"
	"user-service/internal/app/user/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func NewRouteUser(app fiber.Router) {
	serviceUser := container.NewContainerUser()
	HandlerUser := handler.NewHandlerUser(serviceUser)

	app.Post("/login", HandlerUser.Login)
	app.Post("/register", HandlerUser.Register)

	internalRoute := app.Group("/internal").Use(basicauth.New(basicauth.Config{
		Authorizer: func(user, pass string) bool {
			return user == os.Getenv("APP_USERNAME") && pass == os.Getenv("APP_PASSWORD")
		},
	}))

	internalRoute.Get("/users/:id", HandlerUser.GetInternalByID)
}
