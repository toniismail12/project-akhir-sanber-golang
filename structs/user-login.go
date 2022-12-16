package structs

type User_login struct {
	ID       int `json:"id"`
	Username string
	Nama     string
	Role     string
	Token    string
}
