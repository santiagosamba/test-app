package controller

import (
	"github.com/gofiber/fiber/v2"
	"test-app/service"
)

type CharacterController interface {
	GetCharacter(ctx *fiber.Ctx) error
}

type characterController struct {
	characterService service.CharacterService
}

func NewCharacterController(characterService service.CharacterService) CharacterController {
	return &characterController{characterService}
}

func (c *characterController) GetCharacter(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	resp, err := c.characterService.GetCharacter(id)
	if err != nil {
		return err
	}

	return ctx.JSON(resp.Data)
}
