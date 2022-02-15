package create_walet

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type CreateWaletService struct {
	waletRepository infrastructure.WaletRepository
}

func NewCreateWaletService(waletRepo infrastructure.WaletRepository) CreateWaletService {
	return CreateWaletService{
		waletRepository: waletRepo,
	}
}

func (s *CreateWaletService) CreateWalet(ctx context.Context, req CreateWaletRequest, userID int) error {
	errCreate := s.waletRepository.CreateWalet(ctx, RequestMapper(req, userID))
	if errCreate != nil {
		log.Fatal("Service - CreateWalet error : ", errCreate)
	}
	return nil
}
