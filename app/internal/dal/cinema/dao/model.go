package dao

import (
	"Cinema/internal/domain/cinema/model"
	"database/sql"
	"time"
)

type CinemaStorage struct {
	ID        int
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

func (cs *CinemaStorage) ToDomain() model.Cinema {
	return model.Cinema{
		ID:      cs.ID,
		Name:    cs.Name,
		Address: cs.Address,
	}
}
