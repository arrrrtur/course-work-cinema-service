package repository

import (
	"Cinema/internal/domain/ticket/model"
	psql "Cinema/pkg/postgresql"
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type TicketRepositoryInterface interface {
	CreateTicket(ctx context.Context, ticket model.Ticket) error
	GetTicketByID(ctx context.Context, id int) (model.Ticket, error)
	GetTicketsByUserID(ctx context.Context, userId int) ([]model.Ticket, error)
	UpdateTicket(ctx context.Context, id int, updatedTicket model.Ticket) error
	DeleteTicket(ctx context.Context, id int) error
}

type TicketRepository struct {
	qb     squirrel.StatementBuilderType
	client psql.Client
}

func NewTicketRepository(client psql.Client) *TicketRepository {
	return &TicketRepository{
		qb:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		client: client,
	}
}

func (repository *TicketRepository) CreateTicket(ctx context.Context, ticket model.Ticket) error {
	sql, args, err := repository.qb.
		Insert("public.ticket").
		Columns(
			"class",
			"cost",
			"seat",
			"session_id",
			"user_id").
		Values(
			ticket.Class,
			ticket.Cost,
			ticket.Seat,
			ticket.SessionId,
			ticket.UserId,
		).ToSql()
	if err != nil {
		return fmt.Errorf("failed to create ticket: %w", err)
	}

	_, err = repository.client.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (repository *TicketRepository) GetTicketByID(ctx context.Context, id int) (model.Ticket, error) {
	statement := repository.qb.
		Select("class", "cost", "seat", "session_id", "user_id").
		From("public.ticket").
		Where(squirrel.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		return model.Ticket{}, fmt.Errorf("failed to get ticket by ID: %w", err)
	}

	row := repository.client.QueryRow(ctx, query, args...)

	var ticket model.Ticket
	err = row.Scan(&ticket.Class, &ticket.Cost, &ticket.Seat, &ticket.SessionId, &ticket.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Ticket{}, fmt.Errorf("ticket not found")
		}
		return model.Ticket{}, fmt.Errorf("failed to scan row: %w", err)
	}

	return ticket, nil
}

func (repository *TicketRepository) GetTicketsByUserID(ctx context.Context, userId int) ([]model.Ticket, error) {
	statement := repository.qb.
		Select("class", "cost", "seat", "session_id", "user_id").
		From("public.ticket").
		Where(squirrel.Eq{"user_id": userId})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return nil, err
	}

	rows, err := repository.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		return nil, err
	}
	defer rows.Close()

	var tickets []model.Ticket
	for rows.Next() {
		var ticket model.Ticket
		if err = rows.Scan(&ticket.ID, &ticket.UserId, &ticket.Class, &ticket.Seat, &ticket.Cost, &ticket.SessionId); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (repository *TicketRepository) UpdateTicket(ctx context.Context, id int, updatedTicket model.Ticket) error {
	statement := repository.qb.
		Update("public.ticket").
		SetMap(
			map[string]interface{}{
				"class":      updatedTicket.Class,
				"cost":       updatedTicket.Cost,
				"seat":       updatedTicket.Seat,
				"session_id": updatedTicket.SessionId,
				"user_id":    updatedTicket.UserId,
			}).
		Where(squirrel.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		return fmt.Errorf("failed to update ticket: %w", err)
	}

	cmd, err := repository.client.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("no rows updated")
	}

	return nil
}

func (repository *TicketRepository) DeleteTicket(ctx context.Context, id int) error {
	statement := repository.qb.
		Delete("public.ticket").
		Where(squirrel.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		return fmt.Errorf("failed to delete ticket: %w", err)
	}

	cmd, err := repository.client.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("no rows deleted")
	}

	return nil
}
