package repository

import (
	"Cinema/internal/domain/cinema/model"
	"Cinema/pkg/common/logging"
	psql "Cinema/pkg/postgresql"
	"context"
	"github.com/Masterminds/squirrel"
)

type CinemaRepositoryInterface interface {
	Create(ctx context.Context, req model.CreateCinema) error // TODO redeclared create method
	FindAll(ctx context.Context) ([]model.Cinema, error)
	FindById(ctx context.Context, cinemaID int) (model.Cinema, error)
	UpdateBy(ctx context.Context, req model.Cinema) error
	DeleteById(ctx context.Context, cinemaId int) error
}

type CinemaRepository struct {
	qb     squirrel.StatementBuilderType
	client psql.Client
}

func NewCinemaRepository(client psql.Client) *CinemaRepository {
	return &CinemaRepository{
		qb:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		client: client,
	}
}

func (repository *CinemaRepository) Create(ctx context.Context, req model.CreateCinema) error {
	sql, args, err := repository.qb.
		Insert("public.cinema").
		Columns(
			"name",
			"address").
		Values(
			req.Name,
			req.Address,
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

func (repository *CinemaRepository) FindAll(ctx context.Context) ([]model.Cinema, error) {
	statement := repository.qb.
		Select("id", "name", "address").
		From("cinema c")

	// Получение сформированного SQL-запроса и аргументов
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return nil, err
	}

	// Выполнение запроса на выборку
	rows, err := repository.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		logging.L(ctx).Error("pool is closed")
		return nil, err
	}
	defer rows.Close()

	// Обработка результатов выборки
	var cinemas []model.Cinema
	for rows.Next() {
		var cinema model.Cinema
		if err = rows.Scan(&cinema.ID, &cinema.Name, &cinema.Address); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			return nil, err
		}
		cinemas = append(cinemas, cinema)
	}

	return cinemas, nil
}

func (repository *CinemaRepository) FindById(ctx context.Context, cinemaID int) (model.Cinema, error) {
	statement := repository.qb.
		Select("id",
			"name",
			"address").
		From("public.cinema c").
		Where(squirrel.Eq{"id": cinemaID})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return model.Cinema{}, err
	}

	row := repository.client.QueryRow(ctx, query, args...)

	var cinema model.Cinema
	err = row.Scan(&cinema.ID, &cinema.Name, &cinema.Address)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		return model.Cinema{}, err
	}

	return cinema, nil
}

func (repository *CinemaRepository) UpdateBy(ctx context.Context, req model.Cinema) error {
	statement := repository.qb.
		Update("public.cinema p").
		SetMap(
			map[string]interface{}{
				"name":    req.Name,
				"address": req.Address,
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

func (repository *CinemaRepository) DeleteById(ctx context.Context, cinemaId int) error {
	// Формирование запроса на удаление записи
	statement := repository.qb.
		Delete("public.cinema").
		Where(squirrel.Eq{"id": cinemaId})

	// Получение сформированного SQL-запроса и аргументов
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return err
	}

	// Выполнение запроса на удаление
	_, err = repository.client.Exec(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		return err
	}

	return nil
}
