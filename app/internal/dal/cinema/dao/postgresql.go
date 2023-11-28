package dao

import (
	"Cinema/internal/dal/postgres"
	"Cinema/internal/domain/cinema/model"
	psql2 "Cinema/pkg/postgresql"
	"Cinema/pkg/tracing"
	"context"
	"strconv"

	sq "github.com/Masterminds/squirrel"
)

type CinemaDAO struct {
	qb     sq.StatementBuilderType
	client psql2.Client
}

func NewCinemaDAO(client psql2.Client) *CinemaDAO {
	return &CinemaDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

func (repo *CinemaDAO) All(ctx context.Context) ([]model.Cinema, error) {
	all, err := repo.findBy(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]model.Cinema, len(all))
	for i, e := range all {
		resp[i] = e.ToDomain()
	}

	return resp, nil
}

// CreateCinema
func (repo *CinemaDAO) CreateCinema(ctx context.Context, req model.CreateCinema) error {
	sql, args, err := repo.qb.
		Insert(postgres.CinemaTable).
		Columns(
			"id",
			"name",
			"address",
		).
		Values(
			req.ID,
			req.Name,
			req.Address,
		).ToSql()
	if err != nil {
		err = psql2.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
	}

	tracing.SpanEvent(ctx, "Insert Cinema query")
	tracing.TraceVal(ctx, "sql", sql)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := repo.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql2.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return execErr
	}

	if cmd.RowsAffected() == 0 {
		//return dal.ErrNothingInserted
		return nil
	}

	return nil
}

func (repo *CinemaDAO) findBy(ctx context.Context) ([]CinemaStorage, error) {
	statement := repo.qb.
		Select(
			"id",
			"name",
			"address",
			"created_at",
			"updated_at",
		).
		From(postgres.CinemaTable + " c")

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql2.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	tracing.SpanEvent(ctx, "Select Cinema")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		err = psql2.ErrDoQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	defer rows.Close()

	entities := make([]CinemaStorage, rows.CommandTag().RowsAffected())

	for rows.Next() {
		var e CinemaStorage
		if err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Address,
			&e.CreatedAt,
			&e.UpdatedAt,
		); err != nil {
			err = psql2.ErrScan(psql2.ParsePgError(err))
			tracing.Error(ctx, err)

			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}
