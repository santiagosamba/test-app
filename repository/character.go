package repository

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"test-app/defines"
	"test-app/domain"
	"test-app/utils"
)

type CharacterRepository interface {
	GetCharacter(id string) (*domain.Character, error)
}

type characterRepository struct {
	httpClient utils.HTTPClient
}

func NewCharacterRepository(httpClient utils.HTTPClient) CharacterRepository {
	return &characterRepository{httpClient: httpClient}
}

func (c *characterRepository) GetCharacter(id string) (*domain.Character, error) {
	response, err := c.httpClient.Get(fmt.Sprintf("%s%s", defines.DisneyCharactersAPIURL, id))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, defines.ErrGettingCharacterInfo)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, defines.ErrReadingResponse)
	}

	var character domain.Character
	err = json.Unmarshal(body, &character)
	if character.IsEmpty() {
		return nil, fiber.NewError(fiber.StatusNotFound, defines.ErrCharacterInfoNotFound)
	}

	return &character, nil
}
