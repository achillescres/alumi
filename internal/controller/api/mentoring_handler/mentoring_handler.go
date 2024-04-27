package mentoring_handler

import (
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"itamconnect/internal/service/mentoring_serv"
	"net/http"
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

}
