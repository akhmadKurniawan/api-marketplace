package create_product

import (
	"app/application/infrastructure"
	"context"
	"fmt"
	"log"
)

type CreateProductService struct {
	productRepository infrastructure.ProductRepository
	shopRepository    infrastructure.ShopRepository
	sellerRepository  infrastructure.SellerRepository
}

func NewCreateProductService(productRepo infrastructure.ProductRepository, shopRepo infrastructure.ShopRepository, sellerRepo infrastructure.SellerRepository) CreateProductService {
	return CreateProductService{
		productRepository: productRepo,
		shopRepository:    shopRepo,
		sellerRepository:  sellerRepo,
	}
}

func (s *CreateProductService) CreateProduct(ctx context.Context, req CreateProductRequest) error {
	// get seller id
	seller, errGetSeller := s.sellerRepository.GetSellerByUserID(ctx, req.UserID)
	if errGetSeller != nil {
		log.Fatal("Service - GetSellerByUserID error : ", errGetSeller)
	}

	fmt.Println(seller.ID)
	fmt.Println(seller.Shop.SellerID)
	// get shop id
	shop, errGetshop := s.shopRepository.GetShopBySellerID(ctx, seller.Shop.ID)
	fmt.Println(shop.ID)
	fmt.Println(seller.Shop.SellerID)
	if errGetshop != nil {
		log.Fatal("Service - GetshopByUserID error : ", errGetshop)
	}

	errCreate := s.productRepository.CreateProduct(ctx, RequestMapper(req, shop.ID))
	if errCreate != nil {
		log.Fatal("Service - CreateProduct error : ", errCreate)
	}
	return nil
}
