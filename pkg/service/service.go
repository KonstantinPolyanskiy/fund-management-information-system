package service

import (
	internal_types "fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository"
)

type Authorization interface {
	CreateClient(client internal_types.Client) (int, error)
	CreateManager(manager internal_types.SignUp) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Manager interface {
	DeleteById(managerId int) error
	GetById(managerId int) (internal_types.Manager, error)
}
type Client interface {
	Delete(clientId int) error
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
