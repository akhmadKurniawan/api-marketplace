package create_shop

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type CreateShopService struct {
	shopRepository   infrastructure.ShopRepository
	sellerRepository infrastructure.SellerRepository
}

func NewCreateShopService(shopRepo infrastructure.ShopRepository, sellerRepo infrastructure.SellerRepository) CreateShopService {
	return CreateShopService{
		shopRepository:   shopRepo,
		sellerRepository: sellerRepo,
	}
}

func (s *CreateShopService) CreateShop(ctx context.Context, req CreateShopRequest) error {
	//Get seller By Id
	seller, errGetSeller := s.sellerRepository.GetSellerByUserID(ctx, req.UserID)
	if errGetSeller != nil {
		log.Fatal("Service - GetSellerByUserID error : ", errGetSeller)
	}

	errCreate := s.shopRepository.CreateShop(ctx, RequestMapper(req, seller.ID))
	if errCreate != nil {
		log.Fatal("Service - CreateShop error : ", errCreate)
	}
	return nil
}
