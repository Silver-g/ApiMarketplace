package domain

type RegisterUserInternal struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterUserDB struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"hashpassword"`
}
