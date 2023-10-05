package repository

import (
	internal_types "fund-management-information-system/internal-types"
	"fund-management-information-system/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateClient(client internal_types.Client) (int, error)
	CreateManager(manager internal_types.Manager) (int, error)
	User(login, password string) (postgres.User, error)
	Client(login, password string) (internal_types.Client, error)
	Manager(login, password string) (internal_types.Manager, error)
}
type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: postgres.NewAuthPostgres(db)}
}
