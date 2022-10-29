package container

import (
	"exercise-service/internal/app/user/repository"
	"exercise-service/internal/app/user/service"
)

func NewContainerUser() *service.ServiceUser {
	// db := database.NewDatabaseMySQLConnection()
	// userRepository := repository.NewRepositoryDatabaseUser(db)
	userRepository := repository.NewRepositoryMicroserviceUser()
	userService := service.NewServiceUser(userRepository)
	return userService
}
