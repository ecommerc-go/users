package domain

type RegisterUser struct {
	Email    string
	Password string
	Name     string
	Address  string
}

type LoginUser struct {
	Email    string
	Password string
}

type UserProfile struct {
	ID         string
	Email      string
	Name       string
	Address    string
	Created_at string
}

type UpdateProfile struct {
	ID      string
	Name    string
	Address string
}

type Creds struct {
	Login    string
	Password string
	ID       string
}
