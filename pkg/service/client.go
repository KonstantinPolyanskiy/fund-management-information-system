package service

import (
	"fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository"
)

type ClientService struct {
	repo repository.Client
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) DeleteById(clientId int) error {
	return s.repo.DeleteById(clientId)
}

func (s *ClientService) GetById(clientId int) (internal_types.Client, error) {
	return s.repo.GetById(clientId)
}
