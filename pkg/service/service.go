package service

import (
	internal_types "fund-management-information-system/internal-types"
	"fund-management-information-system/pkg/repository"
)

type Authorization interface {
	CreateClient(client internal_types.Client) (int, error)
	CreateManager(manager internal_types.Manager) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
