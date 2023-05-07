package domain

type Reservation struct {
	Start         string
	End           string
	Status        string
	Accommodation Accommodation
	GuestNum      string
	User          User
}

type User struct {
	Name    string
	Surname string
	Email   string
}

type Accommodation struct {
	Name    string
	Id      string
	Address string
}
