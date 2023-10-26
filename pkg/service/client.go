package service

import (
	"context"
	"fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository"
	"github.com/shopspring/decimal"
	"math/rand"
	"time"
)

type ClientService struct {
	repo repository.Client
}

func (s *ClientService) UpdateClient(id int, oldClient internal_types.Client) (internal_types.Client, error) {
	var client internal_types.Client

	if err := s.repo.UpdateClient(id, client); err != nil {
		return client, err
	}
	return s.GetById(id)
}

func (s *ClientService) GetClients(from int) (internal_types.Clients, error) {
	var count int
	clients := make(internal_types.Clients, CountRecords)
	start := from

	for i, _ := range clients {
		if from > CountRecords+start {
			break
		}
		client, err := s.GetById(from)
		if err != nil {
			return nil, err
		}
		if client.Id == 0 {
			break
		}
		clients[i] = client
		from++
		count++
	}

	clientsResult := make(internal_types.Clients, count)

	for i := 0; i < count; i++ {
		clientsResult[i] = clients[i]
	}

	return clientsResult, nil
}

func (s *ClientService) UpdateInvestmentsInfoProcess(ctx context.Context) {
	var client internal_types.Client

	for {
		select {
		case <-ctx.Done():
			return
		default:
			client.InvestmentAmount = investmentAmount(client.InvestmentAmount)

			err := s.repo.UpdateInvestments(client)
			if err != nil {
				ctx.Done()
			}
		}
		time.Sleep(time.Millisecond)
	}
}

func investmentAmount(past decimal.Decimal) decimal.Decimal {
	var next decimal.Decimal

	rand.Seed(int64(uint64(time.Now().UnixNano())))

	random := decimal.NewFromFloat(float64(rand.Intn(4327-1654) + 1654))

	if rand.Float32() < 0.5 {
		next = past.Add(random)
	} else {
		next = past.Sub(random)
	}
	return next
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
