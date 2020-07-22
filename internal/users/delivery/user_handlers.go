package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/pycnick/test_auth_service/internal/users"
)

type UserHandlers struct {
	r *gin.Engine
	uUC *users.UseCase
}

func NewUserHandlers(r *gin.Engine, uUC *users.UseCase) *UserHandlers {
	return &UserHandlers{
		r:   r,
		uUC: uUC,
	}
}

type UserSignUpRequest struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}

func (uH *UserHandlers) SignUp(c *gin.Context) {
	req := &UserSignUpRequest{}

	data, err := c.GetRawData()
	if err != nil {
		
	}
}