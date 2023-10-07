package internal_types

type Client struct {
	Id        int    `json:"-" db:"Id"`
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Phone     string `json:"phone" binding:"required"`
	Login     string `json:"login" binding:"required"`
	Password  string `json:"password" binding:"required"`
	ManagerId *int   `json:"managerId,omitempty"`
}
