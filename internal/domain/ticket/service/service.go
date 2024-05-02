package service

import (
	"Cinema/internal/domain/ticket/model"
	"Cinema/internal/domain/ticket/repository"
	"context"
	"fmt"
)

type TicketServiceInterface interface {
	CreateTicket(ctx context.Context, ticket model.Ticket) error
	GetTicketByID(ctx context.Context, id int) (model.Ticket, error)
	GetTicketsByUserID(ctx context.Context, userId int) ([]model.Ticket, error)
	UpdateTicket(ctx context.Context, id int, updatedTicket model.Ticket) error
	DeleteTicket(ctx context.Context, id int) error
}

type TicketService struct {
	repo repository.TicketRepositoryInterface
}

func NewTicketService(repo repository.TicketRepositoryInterface) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) CreateTicket(ctx context.Context, ticket model.Ticket) error {
	// Add any business logic or validation here before calling the repository method.
	return s.repo.CreateTicket(ctx, ticket)
}

func (s *TicketService) GetTicketByID(ctx context.Context, id int) (model.Ticket, error) {
	// Add any business logic or validation here before calling the repository method.
	if id <= 0 {
		return model.Ticket{}, fmt.Errorf("invalid ticket ID")
	}

	return s.repo.GetTicketByID(ctx, id)
}

func (s *TicketService) GetTicketsByUserID(ctx context.Context, userId int) ([]model.Ticket, error) {
	return s.repo.GetTicketsByUserID(ctx, userId)
}

func (s *TicketService) UpdateTicket(ctx context.Context, id int, updatedTicket model.Ticket) error {
	// Add any business logic or validation here before calling the repository method.
	if id <= 0 {
		return fmt.Errorf("invalid ticket ID")
	}

	return s.repo.UpdateTicket(ctx, id, updatedTicket)
}

func (s *TicketService) DeleteTicket(ctx context.Context, id int) error {
	// Add any business logic or validation here before calling the repository method.
	if id <= 0 {
		return fmt.Errorf("invalid ticket ID")
	}

	return s.repo.DeleteTicket(ctx, id)
}
