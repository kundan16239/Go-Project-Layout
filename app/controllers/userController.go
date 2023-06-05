package controllers

import (
	"context"
	"go-folder-sample/app/models"
	"go-folder-sample/app/services"
	"go-folder-sample/utils/responses"
	"go-folder-sample/utils/validators"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	login := new(models.Login)

	//validate the request body
	if err := ctx.BodyParser(login); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(responses.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: "error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}
	validate := validators.NewValidator()

	//use the validator library to validate required fields
	if err := validate.Struct(login); err != nil {
		// Return, if some fields are not valid.
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"error": validators.ValidatorErrors(err)},
		})
	}
	reqContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := c.UserService.Login(reqContext, login)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}

	splitResponse := strings.Split(resp, ";")
	accessToken := splitResponse[0]
	refreshToken := splitResponse[1]
	return ctx.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"accessToken": accessToken, "refreshToken": refreshToken},
	})
}

// User Status - [0-Inactive, 1-Active, 2-Soft Delete]
func (c *UserController) Register(ctx *fiber.Ctx) error {
	register := new(models.Register)

	//validate the request body
	if err := ctx.BodyParser(register); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(responses.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: "error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}
	validate := validators.NewValidator()

	//use the validator library to validate required fields
	if err := validate.Struct(register); err != nil {
		// Return, if some fields are not valid.
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"error": validators.ValidatorErrors(err)},
		})
	}
	reqContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := c.UserService.Register(reqContext, register)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    resp,
	})
}

// Kyc Staus Code - [0-Progress, 1-Waiting For Driver Confirmation, 2-Pending Due to Not Available at Your Location,3-Success, 4-Failed]
func (c *UserController) UserExist(ctx *fiber.Ctx) error {

	userExist := new(models.UserExist)

	//validate the request body
	if err := ctx.BodyParser(userExist); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(responses.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: "error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}
	validate := validators.NewValidator()

	//use the validator library to validate required fields
	if err := validate.Struct(userExist); err != nil {
		// Return, if some fields are not valid.
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"error": validators.ValidatorErrors(err)},
		})
	}
	reqContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := c.UserService.UserExist(reqContext, userExist)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

func (c *UserController) RefreshToken(ctx *fiber.Ctx) error {
	refreshToken := new(models.RefreshToken)

	//validate the request body
	if err := ctx.BodyParser(refreshToken); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(responses.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: "error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}
	validate := validators.NewValidator()

	//use the validator library to validate required fields
	if err := validate.Struct(refreshToken); err != nil {
		// Return, if some fields are not valid.
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"error": validators.ValidatorErrors(err)},
		})
	}
	reqContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	accessToken, refreshNewToken, err := c.UserService.RefreshToken(reqContext, refreshToken)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"accessToken": accessToken, "refreshToken": refreshNewToken},
	})
}

func (c *UserController) SendEmail(ctx *fiber.Ctx) error {
	return nil
}

func (c *UserController) ForgotPassword(ctx *fiber.Ctx) error {
	return nil
}

func (c *UserController) ResetPassword(ctx *fiber.Ctx) error {
	return nil
}

func (c *UserController) ChangePassword(ctx *fiber.Ctx) error {
	return nil
}
