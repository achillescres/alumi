package httpHandler

import (
	"github.com/gin-gonic/gin"
	"saina-authorization/internal/application/config"
	service "saina-authorization/internal/domain/services"
	"saina-authorization/internal/infrastructure/controller/handler/middleware"
)

type HTTPHandler interface {
	RegisterRouter(group *gin.RouterGroup)
}

type httpHandlerController struct {
	cfg config.HandlerConfig

	middleware  middleware.Middleware
	authService service.PolicyService
}

func NewHttpHandlerController(
	cfg config.HandlerConfig,
	middleware middleware.Middleware,
	authService service.PolicyService,
) HTTPHandler {
	return &httpHandlerController{
		cfg:         cfg,
		middleware:  middleware,
		authService: authService,
	}
}

func (h *httpHandlerController) RegisterRouter(r *gin.RouterGroup) {
	h.registerAuth(r)
}
