package handler

import (
	"context"
	"strconv"
	"user-service/internal/app/domain"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Login(context.Context, domain.LoginParam) (string, error)
	Register(context.Context, domain.RegisterParam) (string, error)
	GetByID(context.Context, int) (domain.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService}
}

func (uh UserHandler) Login(ctx *gin.Context) {
	var param domain.LoginParam
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}
	token, err := uh.userService.Login(ctx.Request.Context(), param)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func (uh UserHandler) Register(ctx *gin.Context) {
	var param domain.RegisterParam
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}

	token, err := uh.userService.Register(ctx.Request.Context(), param)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func (us UserHandler) GetInternalByID(ctx *gin.Context) {
	paramID := ctx.Param("id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid user id",
		})
		return
	}
	user, err := us.userService.GetByID(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, user)
}
