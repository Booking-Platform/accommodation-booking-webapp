package domain

type Reservation struct {
	Id            string
	Start         string
	End           string
	Status        string
	Accommodation Accommodation
	GuestNum      string
	User          User
}

type User struct {
	Id       string
	Name     string
	Surname  string
	Email    string
	Password string
	Role     string
}

type Accommodation struct {
	Name    string
	Id      string
	Address string
}
