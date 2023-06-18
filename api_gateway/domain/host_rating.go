package domain

type HostRating struct {
	Id     string
	Rating string
	Date   string
	Host   User
	Guest  User
}
