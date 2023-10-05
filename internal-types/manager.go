package internal_types

type Manager struct {
	Id       int    `db:"Id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
