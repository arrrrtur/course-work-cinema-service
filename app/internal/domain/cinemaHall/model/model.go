package model

type CinemaHall struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Class    string `json:"class"`
	CinemaId int    `json:"cinema_id"`
}

func NewCinemaHall(
	ID int,
	Name string,
	Capacity int,
	Class string,
	CinemaId int,
) CinemaHall {
	return CinemaHall{
		ID:       ID,
		Name:     Name,
		Capacity: Capacity,
		Class:    Class,
		CinemaId: CinemaId,
	}
}

type CreateCinemaHall struct {
	ID       int    `json:"ID,omitempty"`
	Name     string `json:"name,omitempty"`
	Capacity int    `json:"capacity,omitempty"`
	Class    string `json:"class,omitempty"`
	CinemaId int    `json:"cinemaId,omitempty"`
}

func NewCreateCinemaHall(
	ID int,
	name string,
	Capacity int,
	Class string,
	CinemaId int,
) CreateCinemaHall {
	return CreateCinemaHall{
		ID:       ID,
		Name:     name,
		Capacity: Capacity,
		Class:    Class,
		CinemaId: CinemaId,
	}
}
