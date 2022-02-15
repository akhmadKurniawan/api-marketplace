package create_transaction

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type CreateTransactionHandler struct {
	transactionService CreateTransactionService
}

func NewCreateTransactionHandler(transactionServ CreateTransactionService) CreateTransactionHandler {
	return CreateTransactionHandler{
		transactionService: transactionServ,
	}
}

func (h *CreateTransactionHandler) CreateTransaction(c *gin.Context) {
	req := CreateTransactionRequest{}
	ctx := c.Request.Context()
	acc, _ := c.Get("UserId")
	accId := strconv.FormatInt(acc.(int64), 10)
	userId, _ := strconv.Atoi(accId)

	if err := c.ShouldBind(&req); err != nil {
		log.Fatal("Controller - CreateTransaction error while bind request : ", err)
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Fatal("Controller - CreateTransaction error validation : ", err)
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	errCreate := h.transactionService.CreateTransaction(ctx, req, userId)
	if errCreate != nil {
		log.Fatal("Controller - CreateTransaction error while access service : ", errCreate)
		c.JSON(http.StatusInternalServerError, response.SetMessage(errCreate.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("success", true))
}
