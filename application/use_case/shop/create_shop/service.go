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

func (s *CreateShopService) CreateShop(ctx context.Context, req CreateShopRequest, img string) error {
	//Get seller By Id
	seller, errGetSeller := s.sellerRepository.GetSellerByUserID(ctx, req.UserID)
	if errGetSeller != nil {
		log.Println("Service - GetSellerByUserID error : ", errGetSeller)
		return errGetSeller
	}

	errCreate := s.shopRepository.CreateShop(ctx, RequestMapper(req, seller.ID, img))
	if errCreate != nil {
		log.Println("Service - CreateShop error : ", errCreate)
		return errCreate
	}
	return nil
}
