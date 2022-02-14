package delete_user

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type DeleteUserService struct {
	userRepository infrastructure.UserRepository
	// sellerRepository infrastructure.SellerRepository
}

func NewDeleteUserService(userRepo infrastructure.UserRepository) DeleteUserService {
	return DeleteUserService{
		userRepository: userRepo,
	}
}

func (s *DeleteUserService) DeleteUser(ctx context.Context, id string) error {
	err := s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		log.Println("Service - DeleteUser error : ", err)
		return err
	}
	return nil
}
