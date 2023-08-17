package service

import (
	"github.com/stretchr/testify/mock"
	"test-app/domain"
)

const (
	GetCharacter = "GetCharacter"
)

type CharacterServiceMock struct {
	mock.Mock
}

func (m *CharacterServiceMock) GetCharacter(id string) (*domain.Character, error) {
	args := m.Called(id)
	res := args.Get(0)
	err := args.Get(1)

	if err != nil {
		return nil, err.(error)
	}
	return res.(*domain.Character), nil
}
