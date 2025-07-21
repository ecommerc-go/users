package models

type RegisterRequest struct {
	Email    string
	Password string
	Name     string
	Address  string
}

type LoginUserRequest struct {
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

type UpdateProfileRequest struct {
	ID      string
	Name    string
	Address string
}
