package create_shop

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type CreateShopService struct {
	shopRepository infrastructure.ShopRepository
}

func NewCreateShopService(shopRepo infrastructure.ShopRepository) CreateShopService {
	return CreateShopService{
		shopRepository: shopRepo,
	}
}

func (s *CreateShopService) CreateShop(ctx context.Context, req CreateShopRequest) error {
	errCreate := s.shopRepository.CreateShop(ctx, RequestMapper(req))
	if errCreate != nil {
		log.Fatal("Service - CreateShop error : ", errCreate)
	}
	return nil
}
