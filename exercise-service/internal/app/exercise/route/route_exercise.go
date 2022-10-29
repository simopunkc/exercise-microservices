package route

import (
	"exercise-service/internal/app/container"
	"exercise-service/internal/app/exercise/handler"
	"exercise-service/internal/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRouteExercise(app fiber.Router) {
	serviceExercise := container.NewContainerExercise()
	HandlerExercise := handler.NewHandlerExercise(serviceExercise)

	app.Post("/exercises", middleware.AuthenticationJWT, HandlerExercise.CreateExercise)
	app.Post("/exercises/:id/questions", middleware.AuthenticationJWT, HandlerExercise.CreateQuestion)
	app.Post("/exercises/:id/questions/:qid/answers", middleware.AuthenticationJWT, HandlerExercise.CreateAnswer)

	app.Get("/exercises/:id", middleware.AuthenticationJWT, HandlerExercise.GetExerciseByID)
	app.Get("/exercises/:id/score", middleware.AuthenticationJWT, HandlerExercise.CalculateUserScore)
}
