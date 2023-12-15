package service

import (
	"Cinema/internal/domain/session/model"
	"Cinema/internal/domain/session/repository"
	"context"
	"fmt"
)

type SessionServiceInterface interface {
	CreateSession(ctx context.Context, createSession model.Session) error
	GetSessionByID(ctx context.Context, id int) (model.Session, error)
	GetSessionsByCinemaHallID(ctx context.Context, cinemaHallID int) ([]model.Session, error)
	UpdateSession(ctx context.Context, updatedSession model.Session) error
	DeleteSession(ctx context.Context, id int) error
}

type SessionService struct {
	repository repository.SessionRepositoryInterface
}

func NewSessionService(repository repository.SessionRepositoryInterface) *SessionService {
	return &SessionService{
		repository: repository,
	}
}

func (s *SessionService) CreateSession(ctx context.Context, createSession model.Session) error {
	return s.repository.CreateSession(ctx, createSession)
}

func (s *SessionService) GetSessionsByCinemaHallID(ctx context.Context, cinemaHallID int) ([]model.Session, error) {
	sessions, err := s.repository.GetSessionsByCinemaHallID(ctx, cinemaHallID)
	if err != nil {
		return nil, fmt.Errorf("failed to get sessions by cinema hall ID: %w", err)
	}
	return sessions, nil
}

func (s *SessionService) GetSessionByID(ctx context.Context, id int) (model.Session, error) {
	session, err := s.repository.GetSessionByID(ctx, id)
	if err != nil {
		return model.Session{}, fmt.Errorf("failed to get session by ID: %w", err)
	}
	return session, nil
}

func (s *SessionService) UpdateSession(ctx context.Context, updatedSession model.Session) error {
	err := s.repository.UpdateSession(ctx, updatedSession)
	if err != nil {
		return fmt.Errorf("failed to update session: %w", err)
	}
	return nil
}

func (s *SessionService) DeleteSession(ctx context.Context, id int) error {
	err := s.repository.DeleteSession(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	return nil
}
