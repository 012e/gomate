package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/012e/gomate/controllers"
	"github.com/012e/gomate/models"
	"github.com/012e/gomate/utils/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Authenticate all users accessing routes within the use of this middleware.
// Sets `user_id` using `sessionToken` cookie.
func CookieAuthenticator(c *controllers.DefaultController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("sessionToken")
		if err != nil || strings.TrimSpace(cookie) == "" {
			logrus.Infof("failed with status %s", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
			return
		}

		logrus.Infof("user with cookie %s is in authorization step", cookie)
		var session models.Session
		err = c.DB.First(&session, "token = $1", cookie).Error

		if errors.Is(gorm.ErrRecordNotFound, err) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp.Fail("not logged in"))
			return
		} else if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp.FailUnknown())
			return
		}
		ctx.Set("username", session.Username)
		logrus.Debugf("authencticated user %s", session.Username)
	}
}
