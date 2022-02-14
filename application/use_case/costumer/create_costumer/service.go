package create_costumer

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"log"
	"strconv"
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
	// // get username
	// user, errUser := s.userRepository.GetUsername(ctx, req.Data.Username)
	// if errUser != nil {
	// 	log.Fatal("Service - Login error while access username : ", errUser)
	// }

	user := models.User{}

	errCreate := s.costumerRepository.CreateCostumer(ctx, RequestMapper(req))
	if errCreate != nil {
		log.Fatal("Service - CreateCostumer error : ", errCreate)
	}

	userId := strconv.FormatInt(int64(user.ID), 10)

	//Get user By Id
	_, errGetUser := s.userRepository.GetUserID(ctx, userId)
	if errGetUser != nil {
		log.Fatal("Service - GetUserId error : ", errGetUser)
	}

	return nil
}
