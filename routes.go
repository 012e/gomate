package main

import (
	"github.com/012e/gomate/middlewares"
	"github.com/gin-gonic/gin"
)

func createRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", defaultController.Register)
		auth.POST("/login", defaultController.Login)
	}

	private := r.Group("/")
	private.Use(middlewares.CookieAuthenticator(defaultController))
	// binder must be used before authorizator
	private.Use(middlewares.BindDefaultControllerContexts(defaultController))
	private.Use(middlewares.Authorizator(defaultController))
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


				// TODO: implement
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
}
