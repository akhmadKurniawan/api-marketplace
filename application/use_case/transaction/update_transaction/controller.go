package update_transaction

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type UpdateTransactionHandler struct {
	transactionService UpdateTransactionService
}

func NewUpdateTransactionHandler(transactionServ UpdateTransactionService) UpdateTransactionHandler {
	return UpdateTransactionHandler{
		transactionService: transactionServ,
	}
}

func (h *UpdateTransactionHandler) UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	request := UpdateTransactionRequest{}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&request); !ok {
		log.Println("Controller - UpdateTransaction error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	request.IdVa = id
	res, err := h.transactionService.UpdateTransaction(ctx, request)
	if err != nil {
		c.JSON(400, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(200, SetResponse(res))
}
