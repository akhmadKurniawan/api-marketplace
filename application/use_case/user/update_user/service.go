package update_user

import (
	"app/application/infrastructure"
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UpdateUserService struct {
	userRepository infrastructure.UserRepository
}

func NewUpdateUserService(userRepo infrastructure.UserRepository) UpdateUserService {
	return UpdateUserService{
		userRepository: userRepo,
	}
}

func (s *UpdateUserService) UpdateUser(ctx context.Context, req UpdateUserRequest, id string) (*Response, error) {
	user, err := s.userRepository.GetUserID(ctx, id)
	if err != nil {
		log.Println("Service - Error get user id : ", err)
		return nil, err
	}

	byteDBPass := []byte(user.Password)

	userName, err := s.userRepository.GetUsername(ctx, req.Username)
	fmt.Println("u", userName)
	if err != nil {
		log.Println("Service - UpdateUser error :", err)
		return nil, err
	}

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errHash != nil {
		log.Println("Service - Error hash password : ", errHash)
		return nil, errHash
	}
	if errComparePassword := bcrypt.CompareHashAndPassword(byteDBPass, []byte(req.Password)); errComparePassword != nil {
		log.Println("Service - Password salah : ", errComparePassword)
		return nil, errComparePassword
	}

	_, errUpdate := s.userRepository.UpdateUser(ctx, RequestMapper(req, string(hashedPassword), userName.Username), id)
	if errUpdate != nil {
		log.Println("Service - UpdateUser error :", errUpdate)
		return nil, errUpdate
	}

	res, err := s.userRepository.GetUserID(ctx, id)
	if err != nil {
		log.Println("Service - Get error :", err)
		return nil, err
	}

	return &Response{User: res}, nil

}
