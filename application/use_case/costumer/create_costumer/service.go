package create_costumer

import (
	"app/application/infrastructure"
	"context"
	"fmt"
	"log"
)

type CreateCostumerService struct {
	costumerRepository infrastructure.CostumerRepository
	userRepository     infrastructure.UserRepository
}

func NewCreateCostumerService(costumerRepo infrastructure.CostumerRepository, userRepo infrastructure.UserRepository) CreateCostumerService {
	return CreateCostumerService{
		costumerRepository: costumerRepo,
		userRepository:     userRepo,
	}
}

func (s *CreateCostumerService) CreateCostumer(ctx context.Context, req CreateCostumerRequest) error {

	//Get user By Id
	user, errGetUser := s.userRepository.GetUserID(ctx, req.UserID)
	if errGetUser != nil {
		log.Fatal("Service - GetUserId error : ", errGetUser)
	}
	fmt.Println(user.ID)

	errCreate := s.costumerRepository.CreateCostumer(ctx, RequestMapper(req, user.ID), user.ID)
	if errCreate != nil {
		log.Fatal("Service - CreateCostumer error : ", errCreate)
	}

	return nil
}
