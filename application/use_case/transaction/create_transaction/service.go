package create_transaction

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type CreateTransactionService struct {
	transactionRepository infrastructure.TransactionRepository
}

func NewCreateTransactionService(transactionRepo infrastructure.TransactionRepository) CreateTransactionService {
	return CreateTransactionService{
		transactionRepository: transactionRepo,
	}
}

func (s *CreateTransactionService) CreateTransaction(ctx context.Context, req CreateTransactionRequest, userID int) error {
	errCreate := s.transactionRepository.CreateTransaction(ctx, RequestMapper(req, userID))
	if errCreate != nil {
		log.Fatal("Service - CreateTransaction error : ", errCreate)
	}
	return nil
}
