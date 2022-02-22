package update_transaction

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type UpdateTransactionService struct {
	transactionRepository infrastructure.TransactionRepository
}

func NewUpdateTransactionService(transactionRepo infrastructure.TransactionRepository) UpdateTransactionService {
	return UpdateTransactionService{
		transactionRepository: transactionRepo,
	}
}

func (s *UpdateTransactionService) UpdateTransaction(ctx context.Context, req UpdateTransactionRequest) (*Response, error) {
	trans, errUpdte := s.transactionRepository.UpdateTransaction(ctx, RequestMapper(req, "Success"), req.IdVa)
	if errUpdte != nil {
		log.Fatal("Service - UpdateTransaction error : ", errUpdte)
	}

	return &Response{Trans: trans}, nil
}
