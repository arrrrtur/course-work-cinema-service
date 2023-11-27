package model

type Ticket struct {
	Class     string `json:"class,omitempty"`
	Cost      int    `json:"cost,omitempty"`
	Seat      int    `json:"seat,omitempty"`
	SessionId int    `json:"session_id,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
}
