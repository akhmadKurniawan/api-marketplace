package get_product_shopid

import (
	"app/application/infrastructure"
	"context"
	"errors"
	"log"
)

type ShowProductByShopIDService struct {
	productRepository infrastructure.ProductRepository
}

func NewShowProductByShopIDService(productRepo infrastructure.ProductRepository) ShowProductByShopIDService {
	return ShowProductByShopIDService{
		productRepository: productRepo,
	}
}

func (s *ShowProductByShopIDService) ShowProductByShopID(ctx context.Context, id int) (*ShowProductByShopIDResponseData, error) {
	product, err := s.productRepository.GetProductByShopID(ctx, id)
	if err != nil {
		log.Println("ProductService - ShowProductByShopID error :", err)
		return nil, err
	}

	if err != nil || product.ID == 0 {
		err = errors.New("product not found")
		return nil, err
	}

	res := ShowProductByShopIDResponseData{
		ID:          product.ID,
		ProductType: product.ProductType,
		ShopID:      product.ShopId,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Qty:         product.Qty,
		Image:       product.Image,
	}

	return &res, nil
}
