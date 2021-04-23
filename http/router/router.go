package router

import (
	"management-asset/http/controller"
	"management-asset/service"

	"github.com/gin-gonic/gin"
)

func NewUserRoute(router *gin.Engine, srv service.Services) {
	cnt := controller.Controller{Service: srv}

	router.POST("/regist/", cnt.Registration)
	router.POST("/login/", cnt.Login)

	user := router.Group("/user")
	{
		user.GET("/", cnt.FindAll)
		user.GET("/:id", cnt.FindByID)
		user.PUT("/:id", cnt.Update)
		user.DELETE("/:id", cnt.Delete)
	}
}
