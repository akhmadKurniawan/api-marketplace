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

func (s *CreateCostumerService) CreateCostumer(ctx context.Context, req CreateCostumerRequest) error {
	costumer, _ := s.costumerRepository.GetCostumerByUserId(ctx, req.UserID)
	if req.UserID == costumer.UserID {
		log.Fatal("Service - you already created costumer")
	}

	errCreate := s.costumerRepository.CreateCostumer(ctx, RequestMapper(req))
	if errCreate != nil {
		log.Fatal("Service - CreateCostumer error : ", errCreate)
	}

	return nil
}
