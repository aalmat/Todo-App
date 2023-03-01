package handler

import (
	"github.com/aalmat/todo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var item models.TodoItem
	if err := c.BindJSON(&item); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.services.TodoItem.CreateItem(userId, listId, item)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": itemId,
	})
}

type ItemsList struct {
	Data []models.TodoItem `json:"data"`
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	lists, err := h.services.TodoItem.GetAllItem(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	//fmt.Println(lists)

	c.JSON(http.StatusOK, ItemsList{
		lists,
	})

}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.TodoItem.GetItemById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	//fmt.Println(list)
	c.JSON(http.StatusOK, list)

}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.UpdateItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//fmt.Println(input)

	err = h.services.TodoItem.UpdateItem(userId, itemId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		"ok",
	})

}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoItem.DeleteItem(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, StatusResponse{
		"ok",
	})

}
