package routes

import (
	"github.com/gofiber/fiber/v2"
	"test-app/controller"
)

type Routes interface {
	Install(app *fiber.App)
}

type router struct {
	characterController controller.CharacterController
}

func NewRouter(characterController controller.CharacterController) Routes {
	return &router{characterController}
}

func (r *router) Install(app *fiber.App) {
	app.Get("/character/:id", r.characterController.GetCharacter)
}
