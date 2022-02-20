package create_transaction

import (
	"app/application/infrastructure"
	"context"
	"errors"
	"log"
)

type CreateTransactionService struct {
	transactionRepository infrastructure.TransactionRepository
	productRepository     infrastructure.ProductRepository
	waletRepository       infrastructure.WaletRepository
}

func NewCreateTransactionService(transactionRepo infrastructure.TransactionRepository, productRepo infrastructure.ProductRepository, waletRepo infrastructure.WaletRepository) CreateTransactionService {
	return CreateTransactionService{
		transactionRepository: transactionRepo,
		productRepository:     productRepo,
		waletRepository:       waletRepo,
	}
}

func (s *CreateTransactionService) CreateTransaction(ctx context.Context, req CreateTransactionRequest) error {
	product, errProduct := s.productRepository.GetProductByID(ctx, req.ProductID)
	if errProduct != nil {
		log.Println("Service - CreateTransaction errorProduct : ", errProduct)
		return errProduct
	}

	walet, errWalet := s.waletRepository.GetWaletByUserID(ctx, req.UserID)
	if errWalet != nil {
		log.Println("Service - CreateTransaction errorWalet : ", errWalet)
		return errWalet
	}

	if walet.Saldo < (product.Price * req.TotalProduct) {
		errWalet = errors.New("your saldo not enough")
		return errWalet
	}

	if product.Qty < req.TotalProduct {
		errProduct = errors.New("out of stok")
		return errProduct
	}
	qty := product.Qty - req.TotalProduct
	req.Product.Qty = qty
	_, errUp := s.productRepository.UpdateProdut(ctx, product.ID, req.Product.Qty)
	if errUp != nil {
		log.Println("Service - CreateTransaction errorUpdateProduct : ", errUp)
		return errUp
	}

	req.ProductID = product.ID
	errCreate := s.transactionRepository.CreateTransaction(ctx, RequestMapper(req))
	if errCreate != nil {
		log.Println("Service - CreateTransaction errorCreate : ", errCreate)
		return errCreate
	}

	return nil
}
