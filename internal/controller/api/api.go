package api

import (
	"context"
	"fmt"
	"github.com/achillescres/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"itamconnect/internal/controller/api/auth_handler"
	"itamconnect/internal/service/auth_serv"
)

type API struct {
	Addr string
	Auth *auth_handler.AuthHandler
}

func New(addr string, log *logrus.Entry, authService *auth_serv.Auth) *API {
	return &API{
		Addr: addr,
		Auth: auth_handler.New(authService, log),
	}
}

func (api *API) Run(_ context.Context) error {
	ew := utils.NewErrorWrapper("API - Run")

	r := gin.Default()
	api.Auth.RegisterHandler(r.Group("/auth"))

	err := r.Run(api.Addr)
	return ew(fmt.Errorf("run http sever: %w", err))
}
