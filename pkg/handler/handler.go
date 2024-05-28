package handler

import (
	"github.com/Ko8a/CATEC-Portal/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.POST("/", h.createUserManage)
			users.GET("/", h.getUsers)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUserById)
			users.DELETE("/:id", h.deleteUserById)
			users.GET("/group/:id", h.getUsersByGroupId)
		}

		lesson := api.Group("/lesson")
		{
			lesson.POST("/", h.createLesson)
			lesson.GET("/", h.getAllLessons)
			lesson.GET("/:id", h.getLessonById)
			lesson.PUT("/:id", h.updateLessonById)
			lesson.DELETE("/:id", h.deleteLessonById)
			lesson.GET("/today", h.getDaySchedule)
			lesson.GET("/week", h.getWeekSchedule)
		}

		group := api.Group("/manage")
		{
			group.POST("/group", h.createGroup)
			group.GET("/group", h.getAllGroups)
			group.GET("/group:id", h.getGroupById)
			group.POST("/role", h.createRole)
			group.GET("/role", h.getAllRoles)
			group.GET("/role:id", h.getRoleById)
		}

		schedule := api.Group("/schedule")
		{
			schedule.GET("/range", h.getRangeSchedule)
		}

		account := api.Group("/account")
		{
			account.GET("/", h.getSelfInfo)
			account.PUT("/", h.updateSelfInfo)
			account.DELETE("/", h.deleteSelfInfo)
		}

		marks := api.Group("/marks")
		{
			marks.POST("/", h.createMark)
			marks.PUT("/", h.updateMark)
			marks.GET("/:lesson_id/:user_id", h.getMark)
			marks.GET("/user/:user_id", h.getMarksByUser)
			marks.GET("/lesson/:lesson_id", h.getMarksByLesson)
			marks.GET("/", h.getAllMarks)
		}
	}

	return router
}
