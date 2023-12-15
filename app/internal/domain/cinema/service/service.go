package service

import (
	"Cinema/internal/domain/cinema/model"
	"Cinema/internal/domain/cinema/repository"
	"Cinema/pkg/common/logging"
	"context"
)

type CinemaServiceInterface interface {
	CreateCinema(ctx context.Context, req model.CreateCinema) error
	GetAllCinemas(ctx context.Context) ([]model.Cinema, error)
	GetCinemaByID(ctx context.Context, cinemaID int) (model.Cinema, error)
	UpdateCinema(ctx context.Context, req model.Cinema) error
	DeleteCinema(ctx context.Context, cinemaID int) error
}

type CinemaService struct {
	repository repository.CinemaRepositoryInterface
}

func NewCinemaService(repo repository.CinemaRepositoryInterface) *CinemaService {
	return &CinemaService{
		repository: repo,
	}
}

func (s *CinemaService) CreateCinema(ctx context.Context, req model.CreateCinema) error {
	return s.repository.Create(ctx, req)
}

func (s *CinemaService) GetAllCinemas(ctx context.Context) ([]model.Cinema, error) {
	cinemas, err := s.repository.FindAll(ctx)
	if err != nil {
		logging.L(ctx).Error("aasdsad")
	}
	return cinemas, err
}

func (s *CinemaService) GetCinemaByID(ctx context.Context, cinemaID int) (model.Cinema, error) {
	return s.repository.FindById(ctx, cinemaID)
}

func (s *CinemaService) UpdateCinema(ctx context.Context, req model.Cinema) error {
	return s.repository.UpdateBy(ctx, req)
}

func (s *CinemaService) DeleteCinema(ctx context.Context, cinemaID int) error {
	return s.repository.DeleteById(ctx, cinemaID)
}
