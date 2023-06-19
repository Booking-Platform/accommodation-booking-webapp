package domain

type NewRating struct {
	HostID  string `json:"hostID"`
	GuestID string `json:"guestID"`
	Rating  string `json:"rating"`
}
