package service

import (
	"Cinema/internal/domain/cinema/model"
	"context"
	"github.com/pkg/errors"
)

type repository interface {
	All(ctx context.Context) ([]model.Cinema, error)
	Create(ctx context.Context, req model.CreateCinema) error
}

type CinemaService struct {
	repository repository
}

func NewCinemaService(repository repository) *CinemaService {
	return &CinemaService{repository: repository}
}

func (s *CinemaService) All(ctx context.Context) ([]model.Cinema, error) {
	cinemas, err := s.repository.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.All")
	}

	return cinemas, nil
}

func (s *CinemaService) CreateCinema(ctx context.Context, req model.CreateCinema) (model.Cinema, error) {
	// Логика кэширования или другие дополнительные шаги после создания

	err := s.repository.Create(ctx, req)
	if err != nil {
		return model.Cinema{}, err
	}

	return model.NewCinema(
		req.ID,
		req.Name,
		req.Address,
		// ... другие поля
	), nil
}
