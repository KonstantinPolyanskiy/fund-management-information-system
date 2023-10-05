package internal_types

type Client struct {
	Id       int    `json:"-" db:"Id"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Login    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
