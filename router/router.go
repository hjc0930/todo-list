package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"todo-list/middleware"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "It's work!!!")
		})
	}
	return ginRouter
}
