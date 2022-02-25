package middleware

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Payment string `json:"payment"`
}

func SetMessage(msg string, scss bool, payment string) BaseResponse {
	return BaseResponse{
		Message: msg,
		Success: scss,
		Payment: payment,
	}
}
