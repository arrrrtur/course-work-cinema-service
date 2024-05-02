package repository

import (
	"Cinema/internal/domain/user/model"
	psql "Cinema/pkg/postgresql"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByID(ctx context.Context, id int) (model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, id int, updatedUser model.User) error
	DeleteUser(ctx context.Context, id int) error
}

type UserRepository struct {
	qb     squirrel.StatementBuilderType
	client psql.Client
}

func NewUserRepository(client psql.Client) *UserRepository {
	return &UserRepository{
		qb:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		client: client,
	}
}

func (repository *UserRepository) CreateUser(ctx context.Context, user model.User) error {
	sql, args, err := repository.qb.
		Insert("public.user").
		Columns(
			"first_name",
			"last_name",
			"number",
			"email").
		Values(
			user.FirstName,
			user.LastName,
			user.Number,
			user.Email,
		).ToSql()
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	_, err = repository.client.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (repository *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	statement := repository.qb.
		Select("id", "first_name", "last_name", "number", "email").
		From("public.user u")

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
		return nil, err
	}
	defer rows.Close()

	// Обработка результатов выборки
	var users []model.User
	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.Number, &user.Email, &user.LastName, &user.FirstName); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository *UserRepository) GetUserByID(ctx context.Context, id int) (model.User, error) {
	statement := repository.qb.
		Select("id", "first_name", "last_name", "number", "email").
		From("public.user").
		Where(squirrel.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get user by ID: %w", err)
	}

	row := repository.client.QueryRow(ctx, query, args...)

	var user model.User
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Number, &user.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to scan row: %w", err)
	}

	return user, nil
}

func (repository *UserRepository) UpdateUser(ctx context.Context, id int, updatedUser model.User) error {
	statement := repository.qb.
		Update("public.user").
		SetMap(
			map[string]interface{}{
				"first_name": updatedUser.FirstName,
				"last_name":  updatedUser.LastName,
				"number":     updatedUser.Number,
				"email":      updatedUser.Email,
			}).
		Where(squirrel.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
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

func (repository *UserRepository) DeleteUser(ctx context.Context, id int) error {
	statement := repository.qb.
		Delete("public.user").
		Where(squirrel.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
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
