package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olenka--91/reminder-app/internal/domain"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input domain.User

	if err := ctx.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
