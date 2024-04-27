package httpHandler

import (
	"errors"
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	service "saina-authorization/internal/domain/services"
	"saina-authorization/internal/domain/services/sto"
	"saina-authorization/internal/domain/valueobject"
	"saina-authorization/internal/infrastructure/controller/handler/middleware"
	safeObject "saina-authorization/internal/infrastructure/controller/safeobject"
)

// Register is POST method, that registers users
// It gets the JSON with dto.UserCreate model
func (h *httpHandlerController) Register(c *gin.Context) {
	registerInput := &safeObject.RegisterUserInput{}
	err := c.ShouldBindJSON(registerInput)
	if err != nil {
		err := errors.New("error invalid body")
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, "invalid object format")
		return
	}

	if _, ok := valueobject.UserRoles[registerInput.UserRole]; !ok {
		err := errors.New("error invalid user role")
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, "invalid object format")
		return
	}

	regUserInput := sto.NewRegisterUserInput(registerInput.Login, registerInput.Password, registerInput.FullName, registerInput.Position, registerInput.Email, registerInput.PhoneNumber, registerInput.UserRole, registerInput.AirlineCode)
	id, err := h.authService.RegisterUser(c, regUserInput)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "error registering user")
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *httpHandlerController) Login(c *gin.Context) {
	loginUserInput := sto.LoginUserInput{}
	err := c.ShouldBindJSON(&loginUserInput)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, "invalid object format")
		return
	}

	accessToken, refreshToken, err := h.authService.LoginUser(c, &loginUserInput)
	if errors.Is(err, service.ErrNotFound) {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusBadRequest, err, "invalid login or password")
		return
	}
	if errors.Is(err, service.ErrUnapprovedUser) {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusBadRequest, err, "unapproved user")
		return
	}
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "error couldn't get user")
		return
	}

	tokenCookies := safeObject.NewTokensCoookie(accessToken, refreshToken)
	safeObject.AddCookies(c, *tokenCookies)

	c.JSON(http.StatusOK, "login success")
}

func (h *httpHandlerController) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("Refresh")
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnauthorized, err, "empty refresh token")
		return
	}

	pair, err := h.authService.Refresh(c, refreshToken)
	if errors.Is(err, service.ErrUnapprovedUser) {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnauthorized, err, "unapproved user") // TODO separate cases when it's internal and unauth
		return
	}
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnauthorized, err, "invalid refresh token") // TODO separate cases when it's internal and unauth
		return
	}
	tokenCookies := safeObject.NewTokensCoookie(pair.Access, pair.Refresh)
	safeObject.AddCookies(c, *tokenCookies)

	c.JSON(http.StatusOK, "refresh success")
}

func (h *httpHandlerController) Logout(c *gin.Context) {
	//TODO implement me
	//panic("implement me")
	c.JSON(http.StatusOK, "this doesn't do anything now")
}

func (h *httpHandlerController) Check(c *gin.Context) {
	u, err := middleware.GetUser(c)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnauthorized, err, "empty user") // TODO separate cases when it's internal and unauth
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *httpHandlerController) CheckPermission(c *gin.Context) {
	restrictions, err := middleware.GetRestrictions(c)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "failed create restrictions")
		return
	}
	c.JSON(http.StatusOK, restrictions)
}

func (h *httpHandlerController) registerAuth(r *gin.RouterGroup) {
	r = r.Group("/auth")
	r.GET("/_health", h.Alive)
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/refresh", h.Refresh)
	r.POST("/logout", h.Logout)
	r.GET("/check", h.middleware.UserPolicy, h.Check)
	r.GET("/checkPermission", h.middleware.UserPolicy, h.CheckPermission)
}
