package update_user

import (
	"app/application/infrastructure"
	"context"
	"log"
	"strings"

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
		log.Fatal("Service - Error get user id : ", err)
	}

	byteDBPass := []byte(user.Password)

	userName, _ := s.userRepository.GetUsername(ctx, strings.ToLower(req.Username))
	if user.Username == strings.ToLower(req.Username) {
		//
	} else {
		if strings.ToLower(req.Username) == userName.Username {
			log.Fatal("Service - Username sudah tersedia : ", err)
		}
	}

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errHash != nil {
		log.Fatal("Service - Error hash password : ", errHash)
	}
	if errComparePassword := bcrypt.CompareHashAndPassword(byteDBPass, []byte(req.Password)); errComparePassword != nil {
		log.Fatal("Service - Password salah : ", errComparePassword)
	}

	_, errUpdate := s.userRepository.UpdateUser(ctx, RequestMapper(req, string(hashedPassword), strings.ToLower(req.Username)), id)
	if errUpdate != nil {
		log.Fatal("Service - UpdateUser error :", errUpdate)
	}

	res, err := s.userRepository.GetUserID(ctx, id)
	if err != nil {
		log.Fatal("Service - Get error :", err)
	}

	return &Response{User: res}, nil

}
