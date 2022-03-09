package create_user

import (
	"app/application/infrastructure"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	mailgun "github.com/mailgun/mailgun-go/v4"
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

	// MAIL GUN
	privateAPIKey := "80538cd1f2abea678622efa2106ad1f5-e2e3d8ec-71564881"
	var Domain string = "https://api.mailgun.net/v3/sandbox4c2d556b5cc145a192e64fee2cfcd9c7.mailgun.org"

	mg := mailgun.NewMailgun(Domain, privateAPIKey)

	sender := "sender@example.com"
	subject := "Fancy subject!"
	body := "Hello from Mailgun Go!"
	recipient := req.Email

	fmt.Println(recipient)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	// Mail Gun

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
