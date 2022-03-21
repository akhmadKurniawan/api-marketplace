package get_transaction

import (
	"app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShowTransactionHandler struct {
	transactionService ShowTransactionService
}

func NewShowTransactionHandler(transactionServ ShowTransactionService) ShowTransactionHandler {
	return ShowTransactionHandler{
		transactionService: transactionServ,
	}
}

func (h *ShowTransactionHandler) ShowTransaction(c *gin.Context) {
	res, err := h.transactionService.ShowTransaction(c.Request.Context())
	if err != nil {
		log.Println("TransactionController - ShowTransaction error while access service :", err)
		c.JSON(http.StatusInternalServerError, models.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(res, "Success show transaction", true))
}
