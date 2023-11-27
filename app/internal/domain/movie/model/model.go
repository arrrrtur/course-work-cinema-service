package model

type Movie struct {
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	Duration    int                `json:"duration,omitempty"`
	ReleaseYear string             `json:"release_year,omitempty"`
	Director    string             `json:"director,omitempty"`
	Rating      map[string]float32 `json:"rating,omitempty"`
}
