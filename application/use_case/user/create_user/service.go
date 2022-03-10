package create_user

import (
	"app/application/infrastructure"
	"app/shared"
	"context"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserService struct {
	userRepository     infrastructure.UserRepository
	sellerRepository   infrastructure.SellerRepository
	costumerRepository infrastructure.CostumerRepository
}

func NewCreateUserService(userRepo infrastructure.UserRepository, sellerRepo infrastructure.SellerRepository, costumerRepo infrastructure.CostumerRepository) CreateUserService {
	return CreateUserService{
		userRepository:     userRepo,
		sellerRepository:   sellerRepo,
		costumerRepository: costumerRepo,
	}
}

func (s *CreateUserService) CreateUser(ctx context.Context, req CreateUserRequest) error {
	// Hashing password user
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errHash != nil {
		log.Println("Service - Error hash password ; ", errHash)
		return errHash
	}

	go shared.SendMailgun(shared.Mailgun{
		Sender:    "kurniawan@admin.com",
		Subject:   "app-market",
		Body:      "Please verify your email",
		Recipient: req.Email,
	}) // akan jalan dibelakang layar jadi error di function ini akan di skip

	user, errUser := s.userRepository.GetAllUsername(ctx, req.Username)
	if errUser != nil {
		log.Println("Service - CreateUser error : ", errUser)
		return errUser
	}
	if errUser != nil || user.Username != "" {
		errUser = errors.New("username already exists")
		return errUser
	}

	reqUser := RequestMapper(req, string(hashedPassword), "Inactivated")
	cUser, err := s.userRepository.SignUpUser(ctx, reqUser)
	if err != nil {
		log.Println("Service - CreateUser error : ", err)
		return err
	}

	reqSeller, reqCostumer := RequestMappers(req, cUser.ID)
	if err != nil || cUser.Role < 2 {
		fmt.Println(cUser.Role)
		err = s.costumerRepository.CreateCostumer(ctx, reqCostumer)
		if err != nil {
			log.Println("Service - CreateUser error : ", err)
			return err
		}

		return err
	} else {
		err = s.sellerRepository.CreateSeller(ctx, reqSeller)
		if err != nil {
			log.Println("Service - CreateUser error : ", err)
			return err
		}

		return nil
	}
}
