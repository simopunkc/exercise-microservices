package container

import (
	"exercise-service/internal/app/database"
	"exercise-service/internal/app/exercise/repository"
	"exercise-service/internal/app/exercise/service"
)

func NewContainerExercise() *service.ServiceExercise {
	db := database.NewDatabaseMySQLConnection()
	exerciseRepository := repository.NewRepositoryDatabaseExercise(db)
	exerciseService := service.NewServiceExercise(exerciseRepository)
	return exerciseService
}
