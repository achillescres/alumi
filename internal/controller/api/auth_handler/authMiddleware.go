package auth_handler

import (
	"context"
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"itamconnect/ent"
	"net/http"
)

const authorization = "authorization"

func (ah *AuthHandler) AuthorizationMiddleware(c *gin.Context) {
	jwt, err := c.Cookie("Access")
	if err != nil {
		ah.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, "need Access cookie to check")
		return
	}
	user, err := ah.authServ.Check(c, jwt)
	if err != nil {
		ah.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnauthorized, err, "invalid jwt")
		return
	}

	c.Set(authorization, user)
	c.Next()
}

func GetUser(ctx context.Context) (*ent.User, bool) {
	user, ok := ctx.Value(authorization).(*ent.User)
	if !ok {
		return nil, false
	}
	return user, true
}
