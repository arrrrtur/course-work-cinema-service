package model

type Session struct {
	Date         string `json:"date,omitempty"`
	MovieId      int    `json:"movie_id,omitempty"`
	CinemaHallId int    `json:"cinema_hall_id,omitempty"`
	TicketLeft   int    `json:"ticket_left,omitempty"`
}
