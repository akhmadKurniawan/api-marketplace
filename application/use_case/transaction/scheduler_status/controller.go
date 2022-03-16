package scheduler_status

import (
	"context"
	"fmt"
	"log"
	"time"
)

type UpdateSchedulerHandler struct {
	transactionService UpdateSchedulerService
}

func NewUpdateSchedulerHandler(transactionServ UpdateSchedulerService) UpdateSchedulerHandler {
	return UpdateSchedulerHandler{
		transactionService: transactionServ,
	}
}

func (h *UpdateSchedulerHandler) UpdateScheduler() {
	fmt.Println("Success")
	request := UpdateSchedulerRequest{Status: "failed", UpdatedAt: time.Now()}

	err := h.transactionService.UpdateScheduler(context.Background(), request)
	if err != nil {
		log.Println("Controller - UpdateScheduler error update")
	}
}
