package service

import (
	"test-app/domain"
	"test-app/repository"
)

type CharacterService interface {
	GetCharacter(id string) (*domain.Character, error)
}

type characterService struct {
	characterRepository repository.CharacterRepository
}

func NewCharacterService(characterRepository repository.CharacterRepository) CharacterService {
	return &characterService{characterRepository}
}

func (c *characterService) GetCharacter(id string) (*domain.Character, error) {
	characterInfo, err := c.characterRepository.GetCharacter(id)
	if err != nil {
		return nil, err
	}
	return characterInfo, nil
}
