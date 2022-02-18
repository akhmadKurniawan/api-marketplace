package create_transaction

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type CreateTransactionService struct {
	transactionRepository infrastructure.TransactionRepository
	productRepository     infrastructure.ProductRepository
}

func NewCreateTransactionService(transactionRepo infrastructure.TransactionRepository, productRepo infrastructure.ProductRepository) CreateTransactionService {
	return CreateTransactionService{
		transactionRepository: transactionRepo,
		productRepository:     productRepo,
	}
}

func (s *CreateTransactionService) CreateTransaction(ctx context.Context, req CreateTransactionRequest, userID int) error {
	product, err := s.productRepository.GetProductByID(ctx, req.ProductID)
	if err != nil {
		log.Println("Service - CreateTransaction error : ", err)
		return err
	}

	req.ProductID = product.ID

	errCreate := s.transactionRepository.CreateTransaction(ctx, RequestMapper(req, userID))
	if errCreate != nil {
		log.Println("Service - CreateTransaction error : ", errCreate)
		return errCreate
	}
	return nil
}
