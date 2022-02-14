package delete_seller

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type DeleteSellerService struct {
	sellerRepository infrastructure.SellerRepository
}

func NewDeleteSellerService(sellerRepo infrastructure.SellerRepository) DeleteSellerService {
	return DeleteSellerService{
		sellerRepository: sellerRepo,
	}
}

func (s *DeleteSellerService) DeleteSeller(ctx context.Context, id string) error {
	err := s.sellerRepository.DeleteSeller(ctx, id)
	if err != nil {
		log.Println("Service - DeleteSeller error : ", err)
		return err
	}
	return nil
}
