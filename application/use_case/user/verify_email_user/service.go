package verify_email_user

import (
	"app/application/infrastructure"
	"context"
	"fmt"
	"log"
)

type VerifyEmailUserService struct {
	UserRepository infrastructure.UserRepository
}

func NewVerifyEmailUserService(userRepo infrastructure.UserRepository) VerifyEmailUserService {
	return VerifyEmailUserService{
		UserRepository: userRepo,
	}
}

func (s *VerifyEmailUserService) VerifyEmailUser(ctx context.Context, id string, req VerifyEmailUserRequest) error {
	req.Status = "Activated"
	user, errUser := s.UserRepository.GetStatus(ctx, req.Status)
	if errUser != nil {
		log.Println("Service - VerifyEmailUser error :", errUser)
		return errUser
	}
	// if errUser != nil || user.Status == "Activated" {
	// 	errUser = errors.New("you already activated")
	// 	return errUser
	// }
	fmt.Println(user)

	errVerify := s.UserRepository.VerifyEmailUser(ctx, id, RequestMapper(req))
	if errVerify != nil {
		log.Println("Service - VerifyEmailUser error :", errVerify)
		return errVerify
	}

	return nil
}
