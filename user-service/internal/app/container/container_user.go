package container

import (
	"user-service/internal/app/database"
	"user-service/internal/app/user/repository"
	"user-service/internal/app/user/service"
	"user-service/internal/pkg/util"
)

func NewContainerUser() *service.ServiceUser {
	db := database.NewDatabaseMySQLConnection()
	utilBcrypt := util.NewUtilBcrypt()
	utilJwt := util.NewUtilJwt()
	userRepository := repository.NewRepositoryDatabaseUser(db)
	userService := service.NewServiceUser(userRepository, utilBcrypt, utilJwt)
	return userService
}
