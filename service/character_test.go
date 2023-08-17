package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"test-app/defines"
	"test-app/domain"
	"test-app/repository"
	"testing"
)

func TestGetCharacterError(t *testing.T) {
	repositoryMock := new(repository.CharacterRepositoryMock)
	expectedErr := fiber.NewError(fiber.StatusInternalServerError, "error")
	repositoryMock.On(repository.GetCharacter, mock.Anything).Return(nil, expectedErr)
	service := NewCharacterService(repositoryMock)

	resp, respErr := service.GetCharacter(defines.TestID)

	require.Nil(t, resp)
	require.NotNil(t, respErr)
}

func TestGetCharacterOk(t *testing.T) {
	repositoryMock := new(repository.CharacterRepositoryMock)
	expectedResp := &domain.Character{Data: domain.CharacterData{Name: "Juan", ImageURL: "www.imageurl.com/juan"}}
	repositoryMock.On(repository.GetCharacter, mock.Anything).Return(expectedResp, nil)
	service := NewCharacterService(repositoryMock)

	resp, respErr := service.GetCharacter(defines.TestID)

	require.Nil(t, respErr)
	require.NotNil(t, resp)
	require.Equal(t, resp.Data.Name, "Juan")
}
