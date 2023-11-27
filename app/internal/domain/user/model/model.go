package model

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Number    string `json:"number,omitempty"`
	Email     string `json:"email,omitempty"`
}
