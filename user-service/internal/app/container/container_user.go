package container

import (
	"user-service/internal/app/database"
	"user-service/internal/app/user/repository"
	"user-service/internal/app/user/service"
)

func NewContainerUser() *service.ServiceUser {
	db := database.NewDatabaseMySQLConnection()
	userRepository := repository.NewRepositoryDatabaseUser(db)
	userService := service.NewServiceUser(userRepository)
	return userService
}
