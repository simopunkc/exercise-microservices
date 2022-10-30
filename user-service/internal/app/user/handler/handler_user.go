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
			Hash:  "GMlXoYf29zRo",
			Error: "invalid input",
		})
	}

	service := hu.serviceUser.Login(c.Context(), param)
	if service.Error != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GM4l1oLbAhSP",
			Error: service.Error.Error(),
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash:        "GMuJPADw6w9n",
		RawResponse: service.RawResponse,
	})
}

func (hu HandlerUser) Register(c *fiber.Ctx) error {
	var param domain.RegisterParam
	err := c.BodyParser(&param)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMmBoEMSHwqX",
			Error: "invalid input",
		})
	}

	service := hu.serviceUser.Register(c.Context(), param)
	if service.Error != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMx723OqFFJD",
			Error: service.Error.Error(),
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash:        "GM24dvlTZghU",
		RawResponse: service.RawResponse,
	})
}

func (hu HandlerUser) GetInternalByID(c *fiber.Ctx) error {
	paramID := c.Params("id")
	userID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMZI9Gy7lFuW",
			Error: "invalid user id",
		})
	}

	service := hu.serviceUser.GetByID(c.Context(), userID)
	if service.Error != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMpn1Q0seY4o",
			Error: service.Error.Error(),
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash: "GMSKL7POOhYc",
		User: &service.User,
	})
}
