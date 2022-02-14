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

func (s *CreateSellerService) CreateSeller(ctx context.Context, req CreateSellerRequest) error {
	errCreate := s.sellerRepository.CreateSeller(ctx, RequestMapper(req))
	if errCreate != nil {
		log.Fatal("Service - CreateSeller error : ", errCreate)
	}
	return nil
}
