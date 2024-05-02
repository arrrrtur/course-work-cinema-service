package service

import (
	"Cinema/internal/domain/user/model"
	"Cinema/internal/domain/user/repository"
	"context"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByID(ctx context.Context, id int) (model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, id int, updatedUser model.User) error
	DeleteUser(ctx context.Context, id int) error
}

type UserService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) CreateUser(ctx context.Context, user model.User) error {
	return service.userRepository.CreateUser(ctx, user)
}

func (service *UserService) GetUserByID(ctx context.Context, id int) (model.User, error) {
	return service.userRepository.GetUserByID(ctx, id)
}

func (service *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return service.userRepository.GetAllUsers(ctx)
}

func (service *UserService) UpdateUser(ctx context.Context, id int, updatedUser model.User) error {
	return service.userRepository.UpdateUser(ctx, id, updatedUser)
}

func (service *UserService) DeleteUser(ctx context.Context, id int) error {
	return service.userRepository.DeleteUser(ctx, id)
}
