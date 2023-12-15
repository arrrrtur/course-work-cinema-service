package service

import (
	"Cinema/internal/domain/movie/model"
	"Cinema/internal/domain/movie/repository"
	"context"
	"fmt"
)

type MovieServiceInterface interface {
	CreateMovie(ctx context.Context, movie model.Movie) error
	GetAllMovies(ctx context.Context) ([]model.Movie, error)
	GetMovieByID(ctx context.Context, id int) (model.Movie, error)
	UpdateMovie(ctx context.Context, updatedMovie model.Movie) error
	DeleteMovie(ctx context.Context, id int) error
}

type MovieService struct {
	repository repository.MovieRepositoryInterface
}

func NewMovieService(repository repository.MovieRepositoryInterface) *MovieService {
	return &MovieService{
		repository: repository,
	}
}

func (service *MovieService) CreateMovie(ctx context.Context, movie model.Movie) error {
	err := service.repository.CreateMovie(ctx, movie)
	if err != nil {
		return fmt.Errorf("failed to create movie: %v", err)
	}
	return nil
}

func (service *MovieService) GetAllMovies(ctx context.Context) ([]model.Movie, error) {
	movies, err := service.repository.GetAllMovies(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all movies: %v", err)
	}
	return movies, nil
}

func (service *MovieService) GetMovieByID(ctx context.Context, id int) (model.Movie, error) {
	movie, err := service.repository.GetMovieById(ctx, id)

	if err != nil {
		return model.Movie{}, fmt.Errorf("failed to get movie by ID: %v", err)
	}
	return movie, nil
}

func (service *MovieService) UpdateMovie(ctx context.Context, movie model.Movie) error {
	err := service.repository.UpdateMovie(ctx, movie)

	if err != nil {
		return fmt.Errorf("failed to update movie: %v", err)
	}
	return nil
}

func (service *MovieService) DeleteMovie(ctx context.Context, id int) error {
	err := service.repository.DeleteMovie(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete movie: %v", err)
	}
	return nil
}
