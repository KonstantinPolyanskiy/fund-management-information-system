package internal_types

type ManagerAccount struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
type SignUpClient struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	ManagerId int    `json:"managerId"`
}
