package update_transaction

import (
	"app/application/infrastructure"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	var data Response
	var client = &http.Client{}
	reqXen, err := json.Marshal(req)
	if err != nil {
		log.Println("Service - UpdateTransaction error marshal : ", err)
		return nil, err
	}

	url := fmt.Sprintf("https://api.xendit.co/callback_virtual_accounts/external_id=%v/simulate_payment", req.IdVa)
	reqXendit, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqXen))
	if err != nil {
		log.Println("Service - UpdateTransaction error post xendit : ", err)
		return nil, err
	}

	reqXendit.Header.Add("Content-Type", "application/json")
	reqXendit.SetBasicAuth("xnd_development_RAQZ4fOZldxhwHXIV87sP98nu14as14RGQukAUreIseV8xo32CqmC5Clqorm6bs", "")

	res, err := client.Do(reqXendit)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Status == "" {
		data.Status = "COMPLETED"
	}

	if data.ErrorCode == "INVALID_AMOUNT_ERROR" {
		data.Status = "FAILED"
	}

	data.Trans.IdVa = req.IdVa
	trans, errUpdte := s.transactionRepository.UpdateTransaction(ctx, RequestMapper(req, data.Trans.IdVa, data.Message, data.Status), req.IdVa)
	if errUpdte != nil {
		log.Println("Service - UpdateTransaction error : ", errUpdte)
		return nil, errUpdte
	}

	return &Response{Trans: trans}, nil
}
