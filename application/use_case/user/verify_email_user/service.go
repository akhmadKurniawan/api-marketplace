package verify_email_user

import (
	"app/application/infrastructure"
	"context"
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
	errVerify := s.UserRepository.VerifyEmailUser(ctx, id, RequestMapper(req))
	if errVerify != nil {
		log.Println("Service - VerifyEmailUser error :", errVerify)
		return errVerify
	}

	return nil
}
