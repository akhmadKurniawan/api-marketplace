package create_user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type CreateUserHandler struct {
	userService CreateUserService
}

func NewCreateUserHandler(userService CreateUserService) CreateUserHandler {
	return CreateUserHandler{
		userService: userService,
	}
}

func (h *CreateUserHandler) CreateUser(c *gin.Context) {
	req := CreateUserRequest{}

	ctx := c.Request.Context()

	if err := c.Bind(&req); err != nil {
		log.Println("Controller - CreateUser error while binding request to json : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Println("Controller - CreateUser validation : ", err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(err.Error(), false))
		return
	}

	err := h.userService.CreateUser(ctx, req)
	if err != nil {
		log.Println("Controller - CreateUser error while accessing service : ", err)
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("Success", true))
}
