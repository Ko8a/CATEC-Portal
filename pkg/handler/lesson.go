package handler

import (
	"github.com/Ko8a/CATEC-Portal/structure"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createLesson(c *gin.Context) {
	UserId, err := getUserId(c)
	if err != nil {
		return
	}

	var input structure.Lesson
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Service methods
	id, err := h.services.Lesson.Create(UserId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllLessonsResponse struct {
	Data []structure.LessonInfo `json:"data"`
}

func (h *Handler) getAllLessons(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lessons, err := h.services.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllLessonsResponse{
		Data: lessons,
	})
}

func (h *Handler) getWeekSchedule(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todayLessons, err := h.services.GetWeekLessons(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllLessonsResponse{
		Data: todayLessons,
	})
}

func (h *Handler) getDaySchedule(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todayLessons, err := h.services.GetTodayLessons(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllLessonsResponse{
		Data: todayLessons,
	})
}

func (h *Handler) getLessonById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
	}

	lesson, err := h.services.GetLessonById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (h *Handler) updateLessonById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input structure.UpdateLessonInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateLesson(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteLessonById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
	}

	err = h.services.DeleteLesson(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
