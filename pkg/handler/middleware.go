package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	AuthHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) UserIdentify(c *gin.Context) {
	header := c.GetHeader(AuthHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "header is empty")
		return
	}

	headers := strings.Split(header, " ")
	if len(headers) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "wrong header format")
		return
	}

	userId, err := h.services.ParseToken(headers[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "List not found")
		return 0, errors.New("List not found")
	}

	intId, ok := id.(int)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User id invalid type")
		return 0, errors.New("User id invalid type")
	}

	return intId, nil
}
