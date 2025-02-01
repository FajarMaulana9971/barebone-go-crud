package services

import (
	"barebone-go-crud/src/models/entity"
	"barebone-go-crud/src/repositories"
	"context"
	"errors"
)

type UserService interface {
	GetUserById(ctx context.Context, id int64) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) (int64, error)
	UpdateUser(ctx context.Context, id int64, user *entity.User) (int64, error)
	DeleteUser(ctx context.Context, id int64) error
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repository: repo}
}

func (s *userService) GetUserById(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.GetUserById(ctx, id)
}

func (s *userService) CreateUser(ctx context.Context, user *entity.User) (int64, error) {
	if user.Name == "" || user.Email == "" {
		return 0, errors.New("field name and email is mandatory")
	}
	return s.repository.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, id int64, user *entity.User) (int64, error) {
	if id == 0 {
		return 0, errors.New("id is not valid")
	}

	return s.repository.UpdateUser(ctx, id, user)
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is not valid")
	}

	return s.repository.DeleteUser(ctx, id)
}
