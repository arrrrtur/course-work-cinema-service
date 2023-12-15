package repository

import (
	"Cinema/internal/domain/session/model"
	psql "Cinema/pkg/postgresql"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type SessionRepositoryInterface interface {
	CreateSession(ctx context.Context, session model.Session) error
	GetAllSessions(ctx context.Context) ([]model.Session, error)
	GetSessionByID(ctx context.Context, id int) (model.Session, error)
	GetSessionsByCinemaHallID(ctx context.Context, cinemaHallID int) ([]model.Session, error)
	UpdateSession(ctx context.Context, updatedSession model.Session) error
	DeleteSession(ctx context.Context, id int) error
}

type SessionRepository struct {
	qb     squirrel.StatementBuilderType
	client psql.Client
}

func NewSessionRepository(client psql.Client) *SessionRepository {
	return &SessionRepository{
		qb:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		client: client,
	}
}

func (repository *SessionRepository) CreateSession(ctx context.Context, session model.Session) error {
	sql, args, err := repository.qb.
		Insert("public.session").
		Columns(
			"date",
			"movie_id",
			"cinema_hall_id",
			"ticket_left").
		Values(
			session.Date,
			session.MovieId,
			session.CinemaHallId,
			session.TicketLeft,
		).ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return err
	}

	_, execErr := repository.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		return execErr
	}

	return nil
}

func (repository *SessionRepository) GetAllSessions(ctx context.Context) ([]model.Session, error) {
	statement := repository.qb.
		Select("date", "movie_id", "cinema_hall_id", "ticket_left").
		From("public.session s")

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

	var sessions []model.Session
	for rows.Next() {
		var session model.Session
		if err = rows.Scan(&session.Date, &session.MovieId, &session.CinemaHallId, &session.TicketLeft); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func (repository *SessionRepository) GetSessionByID(ctx context.Context, id int) (model.Session, error) {
	statement := repository.qb.
		Select("date", "movie_id", "cinema_hall_id", "ticket_left").
		From("public.session s").
		Where(squirrel.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return model.Session{}, err
	}

	rows, err := repository.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		return model.Session{}, err
	}

	defer rows.Close()

	if rows.Next() {
		var session model.Session
		if err := rows.Scan(&session.Date, &session.MovieId, &session.CinemaHallId, &session.TicketLeft); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			return model.Session{}, err
		}

		return session, nil
	}

	return model.Session{}, fmt.Errorf("no session found with ID: %d", id)
}

func (repository *SessionRepository) GetSessionsByCinemaHallID(ctx context.Context, cinemaHallID int) ([]model.Session, error) {
	statement := repository.qb.
		Select("id", "date", "movie_id", "cinema_hall_id", "ticket_left").
		From("public.session s").
		Where(squirrel.Eq{"cinema_hall_id": cinemaHallID})

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

	var sessions []model.Session
	for rows.Next() {
		var session model.Session
		if err = rows.Scan(&session.ID, &session.Date, &session.MovieId, &session.CinemaHallId, &session.TicketLeft); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func (repository *SessionRepository) UpdateSession(ctx context.Context, updatedSession model.Session) error {
	statement := repository.qb.
		Update("public.session s").
		SetMap(
			map[string]interface{}{
				"id":             updatedSession.ID,
				"date":           updatedSession.Date,
				"movie_id":       updatedSession.MovieId,
				"cinema_hall_id": updatedSession.CinemaHallId,
				"ticket_left":    updatedSession.TicketLeft,
			}).
		Where(squirrel.Eq{"id": updatedSession.ID})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return err
	}

	_, execErr := repository.client.Exec(ctx, query, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		return execErr
	}

	return nil
}

func (repository *SessionRepository) DeleteSession(ctx context.Context, id int) error {
	statement := repository.qb.
		Delete("public.session").
		Where(squirrel.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return err
	}

	_, execErr := repository.client.Exec(ctx, query, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		return execErr
	}

	return nil
}
