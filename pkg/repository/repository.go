package repository

import (
	internal_types "fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateClient(client internal_types.SignUpClient) (int, error)
	CreateManagerAccount(manager internal_types.ManagerAccount) (int, error)
	GetAccount(login, password string) (postgres.Account, error)
}

type Manager interface {
	DeleteById(managerId int) error
	GetById(managerId int) (internal_types.Manager, error)
	UpdateManager(id int, manager internal_types.Manager) error
	UpdateWorkInfo(manager internal_types.Manager) error
}

type Client interface {
	DeleteById(clientId int) error
	GetById(clientId int) (internal_types.Client, error)
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
