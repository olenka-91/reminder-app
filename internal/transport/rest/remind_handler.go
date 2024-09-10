package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olenka--91/reminder-app/internal/domain"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createRemind(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		logrus.Error(err.Error())
	}

	var input domain.Remind
	if err := ctx.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
	}

	id, err := h.services.Remind.Create(userID, input)
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
	userID, err := getUserId(ctx)
	if err != nil {
		logrus.Error(err.Error())
	}
	//logrus.Debug("userID=", userID)
	res, err := h.services.Remind.GetAll(userID)

	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, getAllRemindsData{
		Data: res,
	})
}

func (h *Handler) getRemindByID(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		logrus.Error(err.Error())
	}
	remindID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponce(ctx, http.StatusBadRequest, "bad id in request")
		return
	}

	res, err := h.services.Remind.GetByID(userID, remindID)

	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) deleteRemind(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		logrus.Error(err.Error())
	}

	remindID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponce(ctx, http.StatusBadRequest, "bad id in request")
		return
	}

	err = h.services.Remind.Delete(userID, remindID)
	if err != nil {
		newErrorResponce(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponce{
		Status: "OK",
	})

}

func (h *Handler) updateRemind(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		logrus.Error(err.Error())
	}
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

	err = h.services.Remind.Update(userID, remindID, input)
	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponce{
		Status: "OK",
	})

}
