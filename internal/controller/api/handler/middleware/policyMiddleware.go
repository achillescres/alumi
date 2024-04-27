package middleware

import (
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"saina-authorization/internal/domain/entity"
	"saina-authorization/internal/domain/valueobject"
)

const restrictionsKey = "restrictions"

func (m *middleware) UserPolicy(c *gin.Context) {
	err := m.Jwt(c)
	if err != nil {
		log.Errorln(err)
		return
	}

	user, err := GetUser(c)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "couldn't get user")
		return
	}

	c.Set(restrictionsKey, *entity.NewRestrictions(user.AirlineCode, valueobject.UserRoles[user.Role]))
	c.Next()
}

func GetRestrictions(c *gin.Context) (*entity.UserInfo, error) {
	user, err := GetUser(c)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	userLogin, err := GetUserLogin(c)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	userInfo := entity.NewUserInfo(*user, userLogin, user.AirlineCode, valueobject.UserRoles[user.Role])
	return userInfo, nil
}

//
//func SetRestrictionForDeveloper(ctx context.Context) context.Context {
//	return context.WithValue(ctx, restrictionsKey, *entity.NewRestrictions(nil, "", valueobject.UserRoleDeveloper))
//}
