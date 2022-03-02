package get_product_shopid

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type ShowProductByShopIDHandler struct {
	ProductByShopIDService ShowProductByShopIDService
}

func NewShowProductByShopIDHandler(productByShopIDServ ShowProductByShopIDService) ShowProductByShopIDHandler {
	return ShowProductByShopIDHandler{
		ProductByShopIDService: productByShopIDServ,
	}
}

func (h *ShowProductByShopIDHandler) ShowProductByShopID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := h.ProductByShopIDService.ShowProductByShopID(c.Request.Context(), id)
	if err != nil {
		log.Println("ProductController - ShowProductByShopID error while access service :", err)
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(res, "Success show product shop", true))
}
