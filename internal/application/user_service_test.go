package application

import (
	"errors"
	"testing"

	"crudWithDB/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestUserServiceCreateOK(t *testing.T) {
	repo := new(UserRepoMock)
	service := NewUserService(repo)

	user := domain.User{
		ID:    "1",
		Name:  "Ana",
		Email: "ana@test.com",
	}

	repo.On("Create", user).
		Return(nil).
		Once()

	err := service.Create(user)

	assert.Equal(t, nil, err)
	repo.AssertExpectations(t)
}

func TestUserServiceCreateError(t *testing.T) {
	repo := new(UserRepoMock)
	service := NewUserService(repo)

	user := domain.User{}
	repo.On("Create", user).
		Return(errors.New("error")).
		Once()

	err := service.Create(user)

	assert.Equal(t, errors.New("error"), err)
	repo.AssertExpectations(t)
}

func TestUserServiceGetByIDOK(t *testing.T) {
	repo := new(UserRepoMock)
	service := NewUserService(repo)

	user := domain.User{
		ID:    "1",
		Name:  "Ana",
		Email: "ana@test.com",
	}

	repo.On("GetByID", "1").
		Return(user, nil).
		Once()

	user, err := service.Get("1")

	assert.Equal(t, nil, err)
	assert.Equal(t, user, domain.User{ID: "1", Name: "Ana", Email: "ana@test.com"})
	repo.AssertExpectations(t)
}

func TestUserServiceGetNotFound(t *testing.T) {
	repo := new(UserRepoMock)
	service := NewUserService(repo)

	repo.On("GetByID", "99").
		Return(domain.User{}, errors.New("not found")).
		Once()

	user, err := service.Get("99")

	assert.Equal(t, domain.User{}, user)
	assert.Equal(t, "not found", err.Error())
	repo.AssertExpectations(t)
}

func TestUserServiceUpdateOK(t *testing.T) {
	repo := new(UserRepoMock)
	service := NewUserService(repo)

	user := domain.User{
		ID:    "1",
		Name:  "Nuevo",
		Email: "nuevo@test.com",
	}

	repo.On("Update", user).
		Return(nil).
		Once()

	err := service.Update(user)

	assert.Equal(t, nil, err)
	repo.AssertExpectations(t)
}

func TestUserServiceUpdateError(t *testing.T) {
	repo := new(UserRepoMock)
	service := NewUserService(repo)

	user := domain.User{}

	repo.On("Update", user).
		Return(errors.New("error")).
		Once()

	err := service.Update(user)

	assert.Equal(t, errors.New("error"), err)
	repo.AssertExpectations(t)
}

func TestUserServiceDeleteOK(t *testing.T) {
	repo := new(UserRepoMock)
	service := NewUserService(repo)

	repo.On("Delete", "1").
		Return(nil).
		Once()

	err := service.Delete("1")

	assert.Equal(t, nil, err)
	repo.AssertExpectations(t)
}

func TestUserServiceDeleteError(t *testing.T) {
	repo := new(UserRepoMock)
	service := NewUserService(repo)

	repo.On("Delete", "1").
		Return(errors.New("error")).
		Once()
	err := service.Delete("1")

	assert.Equal(t, errors.New("error"), err)
	repo.AssertExpectations(t)
}
