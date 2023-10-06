package repository

import (
	internal_types "fund-management-information-system/internal-types"
	"fund-management-information-system/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateClient(client internal_types.Client) (int, error)
	CreateManager(manager internal_types.Manager) (int, error)
	User(login, password string) (postgres.User, string, error)
	Client(login, password string) (internal_types.Client, error)
	Manager(login, password string) (internal_types.Manager, error)
}

type Manager interface {
	Delete(managerId int) error
}
type Client interface {
	Delete(clientId int) error
}
type Repository struct {
	Authorization
	Manager
	Client
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		Manager:       postgres.NewManagerPostgres(db),
		Client:        postgres.NewClientPostgres(db),
	}
}
