package handler

import (
	"github.com/Ko8a/CATEC-Portal/structure"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createMark(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input structure.Mark
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.CreateMark(userId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "mark created"})
}

func (h *Handler) updateMark(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input structure.Mark
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.UpdateMark(userId, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "mark updated"})
}

func (h *Handler) getMark(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lessonId, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lesson_id parameter"})
		return
	}

	mark, err := h.services.GetMark(userId, lessonId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mark)
}

func (h *Handler) getMarksByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id parameter"})
		return
	}

	marks, err := h.services.GetMarksByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, marks)
}

func (h *Handler) getMarksByLesson(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lessonID, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lesson_id parameter"})
		return
	}

	marks, err := h.services.GetMark(userId, lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, marks)
}

func (h *Handler) getAllMarks(c *gin.Context) {
	marks, err := h.services.GetAllMarks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, marks)
}
