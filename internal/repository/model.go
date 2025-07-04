package repository

type RepoRegisterReq struct {
	Email    string
	Password string
	Name     string
	Address  string
}

type RepoLoginReq struct {
	Email    string
	Password string
}

type LoginUserRes struct {
	Access_token string
}

type RepoUserProfile struct {
	ID        string `db:"id"`
	Email     string `db:"email"`
	Name      string `db:"name"`
	Address   string `db:"address"`
	CreatedAt string `db:"created_at"`
}

type UpdateProfileReq struct {
	Id      string
	Name    string
	Address string
}
