package service

import (
	"Cinema/internal/domain/cinemaHall/model"
	"Cinema/internal/domain/cinemaHall/repository"
	"context"
)

type CinemaHallServiceInterface interface {
	CreateCinemaHall(ctx context.Context, req model.CreateCinemaHall) error
	GetAllCinemaHalls(ctx context.Context, cinemaID int) ([]model.CinemaHall, error)
	UpdateCinemaHall(ctx context.Context, req model.CinemaHall) error
	DeleteCinemaHall(ctx context.Context, hallID int) error
	GetCinemaHallByID(ctx context.Context, hallID int) (model.CinemaHall, error)
}

type CinemaHallService struct {
	repository repository.CinemaHallRepositoryInterface
}

func NewCinemaHallService(repo repository.CinemaHallRepositoryInterface) *CinemaHallService {
	return &CinemaHallService{
		repository: repo,
	}
}

func (service *CinemaHallService) GetCinemaHallByID(ctx context.Context, hallID int) (model.CinemaHall, error) {
	return service.repository.GetCinemaHallByID(ctx, hallID)
}

func (service *CinemaHallService) CreateCinemaHall(ctx context.Context, req model.CreateCinemaHall) error {
	return service.repository.CreateCinemaHall(ctx, req)
}

func (service *CinemaHallService) GetAllCinemaHalls(ctx context.Context, cinemaID int) ([]model.CinemaHall, error) {
	return service.repository.GetAllCinemaHalls(ctx, cinemaID)
}

func (service *CinemaHallService) UpdateCinemaHall(ctx context.Context, req model.CinemaHall) error {
	return service.repository.UpdateCinemaHall(ctx, req)
}

func (service *CinemaHallService) DeleteCinemaHall(ctx context.Context, hallID int) error {
	return service.repository.DeleteCinemaHall(ctx, hallID)
}
