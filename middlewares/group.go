package middlewares

import (
	"net/http"

	"github.com/012e/gomate/controllers"
	"github.com/012e/gomate/utils/json"
	"github.com/gin-gonic/gin"
)

func EnsureUserHaveGroup(c *controllers.DefaultController) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !c.PermManager.Context.UserHaveGroup {
			g.AbortWithStatusJSON(http.StatusUnauthorized, json.Fail("user must have a group"))
		}
	}
}

func EnsureUserHaveNoGroup(c *controllers.DefaultController) gin.HandlerFunc {
	return func(g *gin.Context) {
		if c.PermManager.Context.UserHaveGroup {
			g.AbortWithStatusJSON(http.StatusUnauthorized, json.Fail("user must not have a group"))
		}
	}
}
