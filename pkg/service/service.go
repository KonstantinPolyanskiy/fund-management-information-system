package service

import (
	"context"
	"fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository"
)

type Authorization interface {
	CreateClient(client internal_types.SignUpClient) (int, error)
	CreateManagerAccount(manager internal_types.ManagerAccount) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Manager interface {
	DeleteById(managerId int) error
	GetById(managerId int) (internal_types.Manager, error)
	UpdateManager(id int, manager internal_types.Manager) (internal_types.Manager, error)
	GetManagers(from int) (internal_types.Managers, error)
	UpdateWorkInfoProcess(ctx context.Context)
}
type Client interface {
	DeleteById(clientId int) error
	GetById(clientId int) (internal_types.Client, error)
	UpdateClient(id int, OldClient internal_types.Client) (internal_types.Client, error)
	GetClients(from int) (internal_types.Clients, error)
	UpdateInvestmentsInfoProcess(ctx context.Context)
}

type Service struct {
	Authorization
	Manager
	Client
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Manager:       NewManagerService(repos.Manager),
		Client:        NewClientService(repos.Client),
	}
}
