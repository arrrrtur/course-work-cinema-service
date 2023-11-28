package model

type Cinema struct {
	ID      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func NewCinema(
	ID int,
	name string,
	address string,
) Cinema {
	return Cinema{
		ID:      ID,
		Name:    name,
		Address: address,
	}
}

type CreateCinema struct {
	ID      int
	Name    string
	Address string
}

func NewCreateCinema(id int, name string, address string) CreateCinema {
	return CreateCinema{
		ID:      id,
		Name:    name,
		Address: address,
	}
}
