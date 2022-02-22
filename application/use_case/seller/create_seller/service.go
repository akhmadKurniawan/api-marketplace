package create_seller

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type CreateSellerService struct {
	sellerRepository infrastructure.SellerRepository
}

func NewCreateSellerService(sellerRepo infrastructure.SellerRepository) CreateSellerService {
	return CreateSellerService{
		sellerRepository: sellerRepo,
	}
}

func (s *CreateSellerService) CreateSeller(ctx context.Context, req CreateSellerRequest, userID int) error {
	errCreate := s.sellerRepository.CreateSeller(ctx, RequestMapper(req, userID))
	if errCreate != nil {
		log.Println("Service - CreateSeller error : ", errCreate)
		return errCreate
	}
	return nil
}
