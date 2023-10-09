package service

import (
	"fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository"
)

type ManagerService struct {
	repo repository.Manager
}

func NewManagerService(repo repository.Manager) *ManagerService {
	return &ManagerService{repo: repo}
}
func (s *ManagerService) GetById(managerId int) (internal_types.Manager, error) {
	return s.repo.GetById(managerId)
}
func (s *ManagerService) DeleteById(managerId int) error {
	return s.repo.DeleteById(managerId)
}
