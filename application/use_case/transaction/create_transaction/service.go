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
	}

	if errWalet != nil || walet.Saldo < product.Price {
		errWalet = errors.New("your saldo not enough")
		return errWalet
	}
	if errProduct != nil || product.Qty < req.TotalProduct {
		errProduct = errors.New("out of stok")
		return errProduct
	}

	req.ProductID = product.ID

	errCreate := s.transactionRepository.CreateTransaction(ctx, RequestMapper(req))
	if errCreate != nil {
		log.Println("Service - CreateTransaction errorCreate : ", errCreate)
		return errCreate
	}
	return nil
}
