package service

import (
	"context"
	"fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/rand"
	"time"
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
	var count int
	managers := make(internal_types.Managers, CountRecords)

	start := from

	for i, _ := range managers {
		if from > CountRecords+start {
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
		count++
	}

	managersResult := make(internal_types.Managers, count)

	for i := 0; i < count; i++ {
		managersResult[i] = managers[i]
	}

	return managersResult, nil
}
func (s *ManagerService) UpdateWorkInfoProcess(ctx context.Context) {
	var manager internal_types.Manager

	for {
		select {
		case <-ctx.Done():
			return
		default:
			manager.CapitalManagment = capital()
			manager.ProfitPercentDay = percent()

			err := s.repo.UpdateWorkInfo(manager)
			if err != nil {
				ctx.Done()
			}
		}
		time.Sleep(time.Millisecond)
	}
}

func percent() float64 {
	rand.Seed(uint64(time.Now().UnixNano()))
	minPercent := 5.0
	maxPercent := 10.0

	return minPercent + rand.Float64()*(maxPercent-minPercent)
}

func capital() decimal.Decimal {
	rand.Seed(uint64(time.Now().UnixNano()))
	minCapital := decimal.NewFromFloat(10000)
	maxCapital := decimal.NewFromFloat(1000000)

	return minCapital.Add(maxCapital.Sub(minCapital).Mul(decimal.NewFromFloat(rand.Float64())))
}
