package repository

import (
	"Cinema/internal/domain/cinemaHall/model"
	psql "Cinema/pkg/postgresql"
	"context"
	"github.com/Masterminds/squirrel"
)

type CinemaHallRepositoryInterface interface {
	CreateCinemaHall(ctx context.Context, req model.CreateCinemaHall) error
	GetAllCinemaHalls(ctx context.Context, cinemaID int) ([]model.CinemaHall, error)
	GetCinemaHallByID(ctx context.Context, hallID int) (model.CinemaHall, error)
	UpdateCinemaHall(ctx context.Context, req model.CinemaHall) error
	DeleteCinemaHall(ctx context.Context, hallID int) error
}

type CinemaHallRepository struct {
	qb     squirrel.StatementBuilderType
	client psql.Client
}

func NewCinemaHallRepository(client psql.Client) *CinemaHallRepository {
	return &CinemaHallRepository{
		qb:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		client: client,
	}
}

func (repository *CinemaHallRepository) CreateCinemaHall(ctx context.Context, req model.CreateCinemaHall) error {
	sql, args, err := repository.qb.
		Insert("public.cinemaHall").
		Columns(
			"name",
			"capacity",
			"class",
			"id").
		Values(
			req.Name,
			req.Capacity,
			req.Class,
			req.ID,
		).ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)

		return err
	}

	_, execErr := repository.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)

		return err
	}

	return nil
}

func (repository *CinemaHallRepository) GetAllCinemaHalls(ctx context.Context, cinemaID int) ([]model.CinemaHall, error) {
	statement := repository.qb.
		Select("id", "name", "capacity", "class", "cinema_id").
		From("public.cinema_hall").
		Where(squirrel.Eq{"cinema_id": cinemaID})

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

	var cinemaHalls []model.CinemaHall
	for rows.Next() {
		var hall model.CinemaHall
		if err = rows.Scan(&hall.ID, &hall.Name, &hall.Capacity, &hall.Class, &hall.CinemaId); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			return nil, err
		}
		cinemaHalls = append(cinemaHalls, hall)
	}

	return cinemaHalls, nil
}

func (repository *CinemaHallRepository) GetCinemaHallByID(ctx context.Context, hallID int) (model.CinemaHall, error) {
	statement := repository.qb.
		Select("name", "capacity", "class", "id").
		From("public.cinemaHall").
		Where(squirrel.Eq{"id": hallID})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return model.CinemaHall{}, err
	}

	row := repository.client.QueryRow(ctx, query, args...)

	var cinemaHall model.CinemaHall
	err = row.Scan(&cinemaHall.Name, &cinemaHall.Capacity, &cinemaHall.Class, &cinemaHall.ID)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		return model.CinemaHall{}, err
	}

	return cinemaHall, nil
}

func (repository *CinemaHallRepository) UpdateCinemaHall(ctx context.Context, req model.CinemaHall) error {
	statement := repository.qb.
		Update("public.cinemaHall").
		SetMap(
			map[string]interface{}{
				"name":     req.Name,
				"capacity": req.Capacity,
				"class":    req.Class,
				"id":       req.ID,
			}).
		Where(squirrel.Eq{"id": req.ID})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)

		return err
	}

	_, execErr := repository.client.Exec(ctx, query, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		return err
	}

	return nil
}

func (repository *CinemaHallRepository) DeleteCinemaHall(ctx context.Context, hallID int) error {
	statement := repository.qb.
		Delete("public.cinemaHall").
		Where(squirrel.Eq{"id": hallID})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return err
	}

	_, err = repository.client.Exec(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		return err
	}

	return nil
}
