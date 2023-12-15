package model

type Cinema struct {
	ID      int    `json:"id, omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func NewCinema(
	ID int,
	Name string,
	Address string,
) Cinema {
	return Cinema{
		ID:      ID,
		Name:    Name,
		Address: Address,
	}
}

type CreateCinema struct {
	ID      int    `json:"-"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func NewCreateCinema(
	ID int,
	name string,
	address string,
) CreateCinema {
	return CreateCinema{
		ID:      ID,
		Name:    name,
		Address: address,
	}
}
