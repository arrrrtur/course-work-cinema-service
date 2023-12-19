package model

type Movie struct {
	ID          int               `json:"ID"`
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Duration    int               `json:"duration,omitempty"`
	ReleaseYear int               `json:"release_year,omitempty"`
	Director    int               `json:"director,omitempty"`
	Rating      map[string]string `json:"rating,omitempty"`
}
