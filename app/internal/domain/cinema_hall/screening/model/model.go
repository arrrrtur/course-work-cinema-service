package model

type screening struct {
	Id           string `json:"id,omitempty"`
	MovieId      string `json:"movie_id,omitempty"`
	CinemaHallId string `json:"cinema_hall_id,omitempty"`
	DateTime     string `json:"date_time,omitempty"`
}
