package main

import (
	"log"
	"os"
	"user-service/internal/app/database"
	"user-service/internal/app/user/handler"
	"user-service/internal/app/user/repository"
	"user-service/internal/app/user/service"
	"user-service/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	route := gin.Default()

	db := database.NewDatabaseConn()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	route.POST("/login", userHandler.Login)
	route.POST("/register", userHandler.Register)

	internalRoute := route.Group("/internal").Use(middleware.WithBasicAuth())
	{
		internalRoute.GET("/users/:id", userHandler.GetInternalByID)
	}

	route.Run(":" + os.Getenv("APP_PORT"))
}
