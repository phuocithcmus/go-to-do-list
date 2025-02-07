package rest

import (
	"net/http"
	"to_do_project/internal/application/interfaces"
	"to_do_project/internal/interface/dto/request"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(r *gin.Engine, service interfaces.UserService) *UserController {
	controller := &UserController{
		service: service,
	}

	r.POST("/user", controller.CreateUserController)

	return controller
}

func (uc *UserController) CreateUserController(r *gin.Context) {
	var createUserRequest request.CreateUserRequest

	if err := r.Bind(&createUserRequest); err != nil {
		r.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	result, err := uc.service.CreateUser(createUserRequest.UserName, createUserRequest.Password)
	if err != nil {
		r.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	r.JSON(http.StatusCreated, result)
}
