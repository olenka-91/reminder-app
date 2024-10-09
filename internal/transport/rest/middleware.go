package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("UserID not found")
	}

	IntID, ok := id.(int)
	if !ok {
		return 0, errors.New("UserID is of bad type")
	}
	return IntID, nil
}

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponce(c, http.StatusUnauthorized, "token is empty")
		return
	}

	UserID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, UserID)

}
