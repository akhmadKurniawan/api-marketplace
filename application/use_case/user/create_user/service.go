package create_user

import (
	"app/application/infrastructure"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserService struct {
	userRepository infrastructure.UserRepository
}

func NewCreateUserService(userRepo infrastructure.UserRepository) CreateUserService {
	return CreateUserService{
		userRepository: userRepo,
	}
}

func (s *CreateUserService) CreateUser(ctx context.Context, req CreateUserRequest) error {
	// Hashing password user
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errHash != nil {
		log.Fatal("Service - Error hash password ; ", errHash)
	}
	err := s.userRepository.SignUpUser(ctx, RequestMapper(req, string(hashedPassword)))
	if err != nil {
		log.Fatal("Service - CreateUser error : ", err)
	}

	return nil
}
