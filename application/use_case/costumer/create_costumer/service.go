package create_costumer

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type CreateCostumerService struct {
	costumerRepository infrastructure.CostumerRepository
}

func NewCreateCostumerService(costumerRepo infrastructure.CostumerRepository) CreateCostumerService {
	return CreateCostumerService{
		costumerRepository: costumerRepo,
	}
}

func (s *CreateCostumerService) CreateCostumer(ctx context.Context, req CreateCostumerRequest, userId int) error {
	errCreate := s.costumerRepository.CreateCostumer(ctx, RequestMapper(req, userId))
	if errCreate != nil {
		log.Fatal("Service - CreateCostumer error : ", errCreate)
	}

	return nil
}
