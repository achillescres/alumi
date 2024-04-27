package auth_handler

import (
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"itamconnect/internal/service/auth_serv"
	"net/http"
)

type AuthHandler struct {
	authServ *auth_serv.Auth
	log      *logrus.Entry
}

func New(authServ *auth_serv.Auth, log *logrus.Entry) *AuthHandler {
	return &AuthHandler{authServ: authServ, log: log}
}

func (ah *AuthHandler) RegisterHandler(r *gin.RouterGroup) {
	r.POST("/register", ah.Register)
	r.POST("/login", ah.Login)
	r.POST("/check", ah.Check)
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var regInput auth_serv.Register
	err := c.ShouldBindJSON(&regInput)
	if err != nil {
		ah.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, "invalid body")
		return
	}

	c.String(http.StatusOK, "ok")
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var loginInput auth_serv.Login
	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		ah.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, "invalid body")
		return
	}

	jwt, err := ah.authServ.Login(c, loginInput)
	c.String(http.StatusOK, jwt)
}

func (ah *AuthHandler) Check(c *gin.Context) {
	u, _ := GetUser(c)
	u.Edges.Mentor, _ = u.QueryMentor().Only(c)
	u.Edges.Menti, _ = u.QueryMenti().Only(c)
	u.Edges.Skills, _ = u.QuerySkills().All(c)
	u.Edges.RealExperiences, _ = u.QueryRealExperiences().All(c)

	c.JSON(http.StatusOK, u)
}
