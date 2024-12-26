package router

import (
	"todo-list/api"
	"todo-list/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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

		v1.POST("user/register", api.UserRegisterHandler())
	}
	return ginRouter
}
