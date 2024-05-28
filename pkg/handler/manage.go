package handler

import (
	"github.com/Ko8a/CATEC-Portal/structure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createGroup(c *gin.Context) {
	UserId, err := getUserId(c)
	if err != nil {
		return
	}

	var input structure.Group
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Service methods
	id, err := h.services.Manage.CreateGroup(UserId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllGroupsResponse struct {
	Data []structure.Group `json:"data"`
}

func (h *Handler) getAllGroups(c *gin.Context) {
	UserId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	groups, err := h.services.Manage.GetAllGroups(UserId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, getAllGroupsResponse{
		Data: groups,
	})
}

func (h *Handler) getGroupById(c *gin.Context) {

}

func (h *Handler) createRole(c *gin.Context) {
	UserId, err := getUserId(c)
	if err != nil {
		return
	}

	var input structure.Role
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Service methods
	id, err := h.services.Manage.CreateRole(UserId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllRolesResponse struct {
	Data []structure.Role `json:"data"`
}

func (h *Handler) getAllRoles(c *gin.Context) {
	UserId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	roles, err := h.services.Manage.GetAllRoles(UserId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, getAllRolesResponse{
		Data: roles,
	})
}

func (h *Handler) getRoleById(c *gin.Context) {

}
