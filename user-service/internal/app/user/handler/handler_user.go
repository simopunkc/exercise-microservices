package handler

import (
	"context"
	"strconv"
	"user-service/internal/app/domain"

	"github.com/gofiber/fiber/v2"
)

type ServiceUser interface {
	Login(context.Context, domain.LoginParam) domain.Service
	Register(context.Context, domain.RegisterParam) domain.Service
	GetByID(context.Context, int64) domain.Service
}

type HandlerUser struct {
	serviceUser ServiceUser
}

func NewHandlerUser(serviceUser ServiceUser) *HandlerUser {
	return &HandlerUser{serviceUser}
}

func (hu HandlerUser) Login(c *fiber.Ctx) error {
	var param domain.LoginParam
	err := c.BodyParser(&param)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "",
			Error: "invalid input",
		})
	}

	service := hu.serviceUser.Login(c.Context(), param)
	if service.Error != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "",
			Error: service.Error.Error(),
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash:        "",
		RawResponse: service.RawResponse,
	})
}

func (hu HandlerUser) Register(c *fiber.Ctx) error {
	var param domain.RegisterParam
	err := c.BodyParser(&param)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "",
			Error: "invalid input",
		})
	}

	service := hu.serviceUser.Register(c.Context(), param)
	if service.Error != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "",
			Error: service.Error.Error(),
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash:        "",
		RawResponse: service.RawResponse,
	})
}

func (hu HandlerUser) GetInternalByID(c *fiber.Ctx) error {
	paramID := c.Params("id")
	userID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "",
			Error: "invalid user id",
		})
	}

	service := hu.serviceUser.GetByID(c.Context(), userID)
	if service.Error != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "",
			Error: service.Error.Error(),
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash: "",
		User: &service.User,
	})
}
