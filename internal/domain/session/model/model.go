package model

import "github.com/jackc/pgx/v5/pgtype"

type Session struct {
	ID           int         `json:"id"`
	Date         pgtype.Date `json:"date,omitempty"`
	MovieId      int         `json:"movie_id,omitempty"`
	CinemaHallId int         `json:"cinema_hall_id,omitempty"`
	TicketLeft   int         `json:"ticket_left,omitempty"`
}
