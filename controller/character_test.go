package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"test-app/domain"
	"test-app/service"
	"testing"
)

func TestGetCharacterError(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	serviceMock := new(service.CharacterServiceMock)
	expectedErr := fiber.NewError(fiber.StatusInternalServerError, "error")
	serviceMock.On(service.GetCharacter, mock.Anything).Return(nil, expectedErr)
	controller := NewCharacterController(serviceMock)

	respErr := controller.GetCharacter(ctx)

	require.NotNil(t, respErr)
}

func TestGetCharacterOk(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	serviceMock := new(service.CharacterServiceMock)
	expectedRes := &domain.Character{Data: domain.CharacterData{Name: "Juan", ImageURL: "www.imageurl.com/juan"}}
	serviceMock.On(service.GetCharacter, mock.Anything).Return(expectedRes, nil)
	controller := NewCharacterController(serviceMock)

	respErr := controller.GetCharacter(ctx)

	require.Nil(t, respErr)
}
