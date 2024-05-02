package repository

import (
	"Cinema/internal/domain/movie/model"
	psql "Cinema/pkg/postgresql"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
)

type MovieRepositoryInterface interface {
	CreateMovie(ctx context.Context, req model.Movie) error
	GetAllMovies(ctx context.Context) ([]model.Movie, error)
	GetMovieById(ctx context.Context, movieId int) (model.Movie, error)
	UpdateMovie(ctx context.Context, req model.Movie) error
	DeleteMovie(ctx context.Context, movieId int) error
}

type MovieRepository struct {
	qb     squirrel.StatementBuilderType
	client psql.Client
}

func NewMovieRepository(client psql.Client) *MovieRepository {
	return &MovieRepository{
		qb:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		client: client,
	}
}

func (repository *MovieRepository) CreateMovie(ctx context.Context, req model.Movie) error {
	sql, args, err := repository.qb.
		Insert("public.movie").
		Columns(
			"title",
			"description",
			"duration",
			"release_year",
			"director",
			"rating").
		Values(
			req.Title,
			req.Description,
			req.Duration,
			req.ReleaseYear,
			req.Director,
			req.Rating,
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

func (repository *MovieRepository) GetAllMovies(ctx context.Context) ([]model.Movie, error) {
	statement := repository.qb.
		Select("title", "description", "duration", "release_year", "director", "rating").
		From("public.movie m")

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

	var movies []model.Movie
	for rows.Next() {
		var movie model.Movie
		if err = rows.Scan(&movie.Title, &movie.Description, &movie.Duration, &movie.ReleaseYear, &movie.Director, &movie.Rating); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))

			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (repository *MovieRepository) GetMovieById(ctx context.Context, movieId int) (model.Movie, error) {
	statement := repository.qb.
		Select("id", "title", "description", "duration", "release_year", "director", "rating").
		From("public.movie m").
		Where(squirrel.Eq{"id": movieId})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		return model.Movie{}, err
	}

	rows, err := repository.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)

		return model.Movie{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var movie model.Movie
		var hstoreData pgtype.Hstore
		if err = rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Duration, &movie.ReleaseYear, &movie.Director, &hstoreData); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			return model.Movie{}, err
		}
		movie.Rating = hstoreToMap(hstoreData)
		if err != nil {
			log.Fatal(err)
		}
		return movie, nil
	}

	return model.Movie{}, fmt.Errorf("movie not found")
}

func hstoreToMap(hstore pgtype.Hstore) map[string]string {
	result := make(map[string]string)

	for k, v := range hstore {
		result[k] = *v
	}

	return result
}

func (repository *MovieRepository) UpdateMovie(ctx context.Context, req model.Movie) error {
	statement := repository.qb.
		Update("public.movie m").
		SetMap(
			map[string]interface{}{
				"id":           req.ID,
				"title":        req.Title,
				"description":  req.Description,
				"duration":     req.Duration,
				"release_year": req.ReleaseYear,
				"director":     req.Director,
				"rating": func(rating map[string]string) string {
					var hstore string
					for k, v := range rating {
						hstore += fmt.Sprintf(" \"%s\" => %s,", k, v)
					}
					return hstore
				}(req.Rating),
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

func (repository *MovieRepository) DeleteMovie(ctx context.Context, movieId int) error {
	statement := repository.qb.
		Delete("public.movie").
		Where(squirrel.Eq{"id": movieId})

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
