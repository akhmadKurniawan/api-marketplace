package verify_email_user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type VerifyEmailUserHandler struct {
	UserService VerifyEmailUserService
}

func NewVerifyEmailUserHandler(userServ VerifyEmailUserService) VerifyEmailUserHandler {
	return VerifyEmailUserHandler{
		UserService: userServ,
	}
}

func (h *VerifyEmailUserHandler) VerifyEmailUser(c *gin.Context) {
	req := VerifyEmailUserRequest{}
	ctx := c.Request.Context()
	id := c.Param("id")

	if err := c.ShouldBind(&req); err != nil {
		log.Println("Controller - VerifyEmailUser error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Println("Controller - VerifyEmailUser error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	errUpdate := h.UserService.VerifyEmailUser(ctx, id, req)
	if errUpdate != nil {
		log.Println("Controller - VerifyEmailUser error while access service : ", errUpdate)
		c.JSON(500, response.SetMessage(errUpdate.Error(), false))
		return
	}

	c.JSON(http.StatusOK, response.SetMessage("success", true))
}
