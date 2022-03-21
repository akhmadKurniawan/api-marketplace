package get_transaction

import (
	"app/application/infrastructure"
	"context"

	"github.com/rs/zerolog/log"
)

type ShowTransactionService struct {
	transactionRepository infrastructure.TransactionRepository
	// userRepository        infrastructure.UserRepository
}

func NewShowTransactionService(transactionRepo infrastructure.TransactionRepository) ShowTransactionService {
	return ShowTransactionService{
		transactionRepository: transactionRepo,
		// userRepository:        userRepo,
	}
}

func (s *ShowTransactionService) ShowTransaction(ctx context.Context) (*Response, error) {
	transaction, err := s.transactionRepository.GetTransactions(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Service - ShowTransaction error")
		return nil, err
	}

	return &Response{Transaction: transaction}, nil
}
