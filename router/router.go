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
	router := ginRouter.Group("")
	{
		router.GET("ping", func(context *gin.Context) {
			context.JSON(200, "It's work!!!")
		})
		// user
		router.POST("user/register", api.UserRegisterHandler())
		router.POST("user/login", api.UserLoginHandler())
		// task router
		authed := router.Group("/task")
		authed.Use(middleware.Auth())
		{
			authed.GET("/list", api.ListTaskHandler())
			authed.GET("/details", api.DetailsTaskHandler())
			authed.GET("/search", api.SearchTaskHandler())
			authed.POST("", api.CreateTaskHandler())
			authed.PUT("", api.UpdateTaskHandler())
			authed.DELETE("", api.DeleteTaskHandler())
		}
	}

	return ginRouter
}
