package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/pycnick/test_auth_service/internal/models"
	"github.com/pycnick/test_auth_service/internal/users"
	"net/http"
)

type UserHandlers struct {
	uUC users.UseCase
}

func NewUserHandlers(r *gin.RouterGroup, uUC users.UseCase) *UserHandlers {
	uH := &UserHandlers{
		uUC: uUC,
	}

	r.POST("/register", uH.SignUp)

	return uH
}

type UserSignUpRequest struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}

func (uH *UserHandlers) SignUp(c *gin.Context) {
	req := &UserSignUpRequest{}

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: BadRequestError.Error(),
		})
		return
	}

	user, err := models.NewUser(req.Login, req.Email, req.Password, req.Phone)

	if err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}

	err = uH.uUC.SignUp(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Message{
		Message: "success",
	})
}