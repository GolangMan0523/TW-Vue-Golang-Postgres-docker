package action

import (
	// "errors"

	// "github.com/HotPotatoC/twitter-clone/internal/common/config"
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/auth/service"
	// "github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/gofiber/fiber/v2"
)

type resetAction struct {
	service service.ResetService
}

func NewResetAction(service service.ResetService) module.Action {
	return resetAction{service: service}
}

func (a resetAction) Execute(c *fiber.Ctx) error {
	var input service.ResetInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := input.Validate(); errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	resetToken, err := a.service.Execute(input)
	if err != nil {
		switch {
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "There was a problem on our side",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "true",
		"token": resetToken,
	})
}
