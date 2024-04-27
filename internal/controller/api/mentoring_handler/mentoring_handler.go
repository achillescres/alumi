package mentoring_handler

import (
	"errors"
	"fmt"
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"itamconnect/internal/controller/api/auth_handler"
	"itamconnect/internal/domain/valueobject"
	"itamconnect/internal/service/mentoring_serv"
	"net/http"
	"strconv"
)

type MentoringHandler struct {
	log *logrus.Entry
	ms  mentoring_serv.MentoringService
}

func (mh *MentoringHandler) RegisterHandler(r *gin.RouterGroup) {
	r.POST("/mentorsForMe", mh.GetMentorsForMe)
}

func (mh *MentoringHandler) GetMentorsForMe(c *gin.Context) {
	var searchMentorsOpts mentoring_serv.SearchMentorsOpts
	err := c.ShouldBindJSON(&searchMentorsOpts)
	mentors, err := mh.ms.GetMentors(c, searchMentorsOpts)
	if err != nil {
		mh.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, mentors)
}

func (mh *MentoringHandler) Request(c *gin.Context) {
	user, _ := auth_handler.GetUser(c)
	if user.Type != valueobject.UserTypeMenti {
		err := errors.New("only menti can request to mentors")
		mh.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnauthorized, err, err.Error())
		return
	}

	mentorIDRaw, ok := c.GetQuery("mentorID")
	if !ok {
		mh.log.Errorln(ginresponse.MissingRequiredQuery(c, "mentorID"))
		return
	}
	mentorID, err := strconv.Atoi(mentorIDRaw)
	if err != nil || mentorID <= 0 {
		err := errors.New("invalid mentorID")
		mh.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, err.Error())
		return
	}

	err = mh.ms.Request(c, user.ID, mentorID)
	if err != nil {
		mh.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.String(http.StatusOK, "ok")
}

func (mh *MentoringHandler) SolveRequest(c *gin.Context) {
	user, _ := auth_handler.GetUser(c)
	if user.Type != valueobject.UserTypeMentor {
		err := errors.New("only mentors can accept requests")
		mh.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnauthorized, err, err.Error())
		return
	}

	matchIDRaw, ok := c.GetQuery("matchID")
	if !ok {
		mh.log.Errorln(ginresponse.MissingRequiredQuery(c, "matchID"))
		return
	}
	matchID, err := strconv.Atoi(matchIDRaw)
	if err != nil || matchID <= 0 {
		err := errors.New("invalid requestID")
		mh.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, err.Error())
		return
	}
	acceptRaw, ok := c.GetQuery("accept")
	if !ok {
		mh.log.Errorln(ginresponse.MissingRequiredQuery(c, "accept"))
		return
	}
	accept, err := strconv.ParseBool(acceptRaw)
	if err != nil {
		err := fmt.Errorf("invalid accept, must be bool: %w", err)
		mh.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, err.Error())
		return
	}
	err = mh.ms.SolveMatch(c, user.ID, matchID, accept)
	if err != nil {
		mh.log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.String(http.StatusOK, "ok")
}
