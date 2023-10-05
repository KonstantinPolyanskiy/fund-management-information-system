package service

import "fund-management-information-system/pkg/repository"

type ManagerService struct {
	repo repository.Manager
}

func NewManagerService(repo repository.Manager) *ManagerService {
	return &ManagerService{repo: repo}
}

func (s *ManagerService) Delete(managerId int) error {
	return s.repo.Delete(managerId)
}
