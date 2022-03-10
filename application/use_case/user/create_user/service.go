package create_user

import (
	"app/application/infrastructure"
	"app/shared"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

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
	UserId, err := s.userRepository.SignUpUser(ctx, reqUser)
	if err != nil {
		log.Println("Service - CreateUser error : ", err)
		return err
	}

	t := time.Now()
	ti := t.Format("20060102150405")
	id := strconv.Itoa(int(UserId.ID))
	combineString := ti + ":" + id

	message := fmt.Sprintf("Please verify your email\nclick this link for verify: localhost:5000/api/v1/users/active/%s", combineString)
	go shared.SendMailgun(shared.Mailgun{
		Sender:    "kurniawan@admin.com",
		Subject:   "app-market",
		Body:      message,
		Recipient: req.Email,
	}) // akan jalan dibelakang layar jadi error di function ini akan di skip

	reqSeller, reqCostumer := RequestMappers(req, UserId.ID)
	if err != nil || UserId.Role < 2 {
		fmt.Println(UserId.Role)
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
