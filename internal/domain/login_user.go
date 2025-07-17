package domain

type LoginUserInternal struct {
	Id           int    `json:"id"`
	Username     string `json:"login"`
	Password     string `json:"password"`
	PasswordHash string `json:"hashpassword"`
}
type LoginUserDB struct {
	Username string `json:"login"`
}
type LoginUserResponseDb struct {
	Id           int    `json:"id"`
	PasswordHash string `json:"hashpassword"`
}
