package service

import (
	"fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository"
)

type ManagerService struct {
	repo repository.Manager
}

const CountRecords = 50

func NewManagerService(repo repository.Manager) *ManagerService {
	return &ManagerService{repo: repo}
}
func (s *ManagerService) GetById(managerId int) (internal_types.Manager, error) {
	return s.repo.GetById(managerId)
}
func (s *ManagerService) DeleteById(managerId int) error {
	return s.repo.DeleteById(managerId)
}
func (s *ManagerService) UpdateManager(id int, oldManager internal_types.Manager) (internal_types.Manager, error) {
	var manager internal_types.Manager
	if err := s.repo.UpdateManager(id, oldManager); err != nil {
		return manager, err
	}
	return s.GetById(id)
}
func (s *ManagerService) GetManagers(from int) (internal_types.Managers, error) {
	managers := make(internal_types.Managers, CountRecords)

	idk := from

	for i, _ := range managers {
		if from > CountRecords+idk {
			break
		}
		manager, err := s.GetById(from)
		if err != nil {
			return nil, err
		}
		if manager.Id == 0 {
			break
		}
		managers[i] = manager
		from++
	}

	return managers, nil
}
