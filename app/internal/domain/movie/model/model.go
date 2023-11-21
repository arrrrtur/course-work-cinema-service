package model

type Movie struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
	Duration    string `json:"duration"`
	ReleaseYear string `json:"release_Year"`
}
