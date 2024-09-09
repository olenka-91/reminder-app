package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olenka--91/reminder-app/internal/domain"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createRemind(ctx *gin.Context) {
	UserID, err := getUserId(ctx)
	if err != nil {
		logrus.Error(err.Error())
	}

	var input domain.Remind
	if err := ctx.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
	}

	id, err := h.services.Remind.Create(UserID, input)
	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllRemindsData struct {
	Data []domain.Remind `json:"data"`
}

func (h *Handler) getAllReminds(ctx *gin.Context) {

	res, err := h.services.Remind.GetAll(0)

	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, getAllRemindsData{
		Data: res,
	})
}

func (h *Handler) getRemindByID(ctx *gin.Context) {

	remindID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponce(ctx, http.StatusBadRequest, "bad id in request")
		return
	}

	res, err := h.services.Remind.GetByID(0, remindID)

	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) deleteRemind(ctx *gin.Context) {
	remindID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponce(ctx, http.StatusBadRequest, "bad id in request")
		return
	}

	err = h.services.Remind.Delete(0, remindID)
	if err != nil {
		newErrorResponce(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponce{
		Status: "OK",
	})

}

func (h *Handler) updateRemind(ctx *gin.Context) {
	remindID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponce(ctx, http.StatusBadRequest, "bad id in request")
		return
	}

	var input domain.RemindUpdateInput
	if err = ctx.BindJSON(&input); err != nil {
		newErrorResponce(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Remind.Update(0, remindID, input)
	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponce{
		Status: "OK",
	})

}
