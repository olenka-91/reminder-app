package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olenka--91/reminder-app/internal/domain"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createRemind(ctx *gin.Context) {
	var input domain.Remind

	if err := ctx.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
	}
	id, err := h.services.Remind.Create(0, input)

	if err != nil {
		logrus.Error(err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllReminds(ctx *gin.Context) {
}
func (h *Handler) getRemindByID(ctx *gin.Context) {
}
func (h *Handler) deleteRemind(ctx *gin.Context) {
}
func (h *Handler) updateRemind(ctx *gin.Context) {
}
