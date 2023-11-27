package model

type CinemaHall struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Class    string `json:"class"`
	CinemaId int    `json:"cinema_id"`
}
