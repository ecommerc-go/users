package service

type RegisterUserRequest struct {
	Email    string
	Password string
	Name     string
	Address  string
}

type RegisterUserResponse struct {
	Id string
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type LoginUserResponse struct {
	Access_token string
}

type GetProfileRequest struct {
	ID string
}

type UserProfile struct {
	Id         string
	Email      string
	Name       string
	Address    string
	Created_at string
}

type GetProfileResponse struct {
	Profile UserProfile
}

type UpdateProfileRequest struct {
	Id      string
	Name    string
	Address string
}
