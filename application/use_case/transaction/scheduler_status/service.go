package scheduler_status

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type UpdateSchedulerService struct {
	transactionRepository infrastructure.TransactionRepository
}

func NewUpdateSchedulerService(transactionRepo infrastructure.TransactionRepository) UpdateSchedulerService {
	return UpdateSchedulerService{
		transactionRepository: transactionRepo,
	}
}

func (s *UpdateSchedulerService) UpdateScheduler(ctx context.Context, req UpdateSchedulerRequest) error {
	errUpdate := s.transactionRepository.UpdateScheduler(ctx, RequestMapper(req))
	if errUpdate != nil {
		log.Println("Service - UpdateScheduler error :", errUpdate)
		return errUpdate
	}
	return nil
}
