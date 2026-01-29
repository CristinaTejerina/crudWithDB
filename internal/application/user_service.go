package application

import "crudWithDB/internal/domain"

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user domain.User) error {
	return s.repo.Create(user)
}

func (s *UserService) Get(id string) (domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Update(user domain.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}
