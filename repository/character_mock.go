package repository

import (
	"github.com/stretchr/testify/mock"
	"test-app/domain"
)

const (
	GetCharacter = "GetCharacter"
)

type CharacterRepositoryMock struct {
	mock.Mock
}

func (m *CharacterRepositoryMock) GetCharacter(id string) (*domain.Character, error) {
	args := m.Called(id)
	res := args.Get(0)
	err := args.Get(1)

	if err != nil {
		return nil, err.(error)
	}
	return res.(*domain.Character), nil
}
