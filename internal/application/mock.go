package application

import (
	"crudWithDB/internal/domain"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) Create(u domain.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *UserRepoMock) GetByID(id string) (domain.User, error) {
	args := m.Called(id)
	user, _ := args.Get(0).(domain.User)
	return user, args.Error(1)
}

func (m *UserRepoMock) Update(u domain.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *UserRepoMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
