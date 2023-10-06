package service

import "fund-management-information-system/pkg/repository"

type ClientService struct {
	repo repository.Client
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) Delete(clientId int) error {
	return s.repo.Delete(clientId)
}
