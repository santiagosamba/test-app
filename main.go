package main

import (
	"github.com/gofiber/fiber/v2"
	"test-app/controller"
	"test-app/repository"
	"test-app/routes"
	"test-app/service"
	"test-app/utils"
)

func main() {
	app := fiber.New()

	httpClient := utils.NewHTTPClient()
	characterRepository := repository.NewCharacterRepository(httpClient)
	characterService := service.NewCharacterService(characterRepository)
	characterController := controller.NewCharacterController(characterService)

	router := routes.NewRouter(characterController)
	router.Install(app)

	app.Listen(":3000")
}
