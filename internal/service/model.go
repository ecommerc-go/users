package service

type SvcRegisterRequest struct {
	Email    string
	Password string
	Name     string
	Address  string
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type LoginUserResponse struct {
	Access_token string
}

type SvcUserProfile struct {
	Id         string
	Email      string
	Name       string
	Address    string
	Created_at string
}

type UpdateProfileRequest struct {
	Id      string
	Name    string
	Address string
}
