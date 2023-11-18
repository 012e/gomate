package middlewares

import (
	"net/http"

	"github.com/012e/gomate/controllers"
	"github.com/012e/gomate/utils/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Authorizator(c *controllers.DefaultController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logrus.Infof("user with id %s is in authorization stage", c.Username)

		// begin permission checking
		ok, err := c.PermManager.Enforce(c.Username, ctx.Request.URL.Path, ctx.Request.Method)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp.Fail("failed to validate user: "+err.Error()))
			return
		}
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, resp.Fail("user unauthorized for "+ctx.Request.URL.Path))
			return
		}
		logrus.Infof("user %s successfully got into %s route", c.Username, ctx.Request.URL.Path)
	}
}
