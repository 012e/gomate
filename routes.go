package main

import (
	"net/http"

	"github.com/012e/gomate/middlewares"
	"github.com/012e/gomate/utils/resp"
	"github.com/gin-gonic/gin"

	docs "github.com/012e/gomate/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func createRoutes(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/", defaultController.Hello)
	r.GET("/todo/:id", defaultController.GetTodo)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	auth := r.Group("/auth")
	{
		auth.POST("/register", defaultController.Register)
		auth.POST("/login", defaultController.Login)
	}

	private := r.Group("/")
	private.Use(middlewares.CookieAuthenticator(defaultController))
	// binder must be used before authorizator
	private.Use(middlewares.BindDefaultControllerContexts(defaultController))
	// private.Use(middlewares.Authorizator(defaultController))
	private.GET("/hello", defaultController.Hello)
	{
		group := private.Group("/group")
		{
			groupNotRequired := group.Group("/")
			groupNotRequired.Use(middlewares.EnsureUserHaveNoGroup(defaultController))
			{
				groupNotRequired.POST("/new", defaultController.CreateGroup)
				groupNotRequired.GET("/join/:code", defaultController.JoinGroup)
			}

			groupRequired := group.Group("/")
			groupRequired.Use(middlewares.EnsureUserHaveGroup(defaultController))
			{
				groupRequired.GET("/:id", defaultController.LeaveGroup)

				// user only have a single group, no need for id param
				groupRequired.GET("/leave", defaultController.LeaveGroup)
				groupRequired.POST("/join/new", defaultController.CreateJoinCode)

				// TODO
				// groupRequired.DELETE("/:code", defaultController.DeleteGroup)
				// groupRequired.GET("/all", defaultController.CreateJoinCode)
			}
		}

		todo := private.Group("/todo")
		todo.Use(middlewares.EnsureUserHaveGroup(defaultController))
		{
			todo.POST("/new", defaultController.CreateTodo)
			// todo.GET("/:id", defaultController.GetTodo)
			// todo.DELETE("/:id", defaultController.GetTodo)
			// todo.PATCH("/:id", defaultController.GetTodo)
		}
	}
	r.NoRoute(func(g *gin.Context) { g.JSON(http.StatusNotFound, resp.Fail("route doesn't exist")) })
}
