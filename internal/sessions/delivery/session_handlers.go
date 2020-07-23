package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/pycnick/test_auth_service/internal/sessions"
	"log"
	"net/http"
)

type SessionHandlers struct {
	sUC sessions.UseCase
}

func NewSessionHandlers(r *gin.RouterGroup, sUC sessions.UseCase) *SessionHandlers {
	sH := &SessionHandlers{
		sUC: sUC,
	}

	r.POST("/login", sH.SignIn)

	return sH
}

type SessionSignInRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func (sH *SessionHandlers) SignIn(c *gin.Context) {
	req := &SessionSignInRequest{}

	if err := c.Bind(req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Error{
			Error: BadRequestError.Error(),
		})
		return
	}

	reqCookie, err := c.Cookie("Test_auth_service")
	if err == nil {
		if _, err := sH.sUC.GetSession(reqCookie); err == nil {
			c.JSON(http.StatusConflict, Message{
				Message: AlreadyAuthenticatedError.Error(),
			})
			return
		}
	}

	session, err := sH.sUC.SignIn(req.Login, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}

	cookie := session.GetCookie()

	c.SetCookie(cookie.Name,
		cookie.Value,
		cookie.MaxAge,
		cookie.Path,
		cookie.Domain,
		cookie.Secure,
		cookie.HttpOnly)

	c.JSON(http.StatusOK, Message{
		Message: "success",
	})
}