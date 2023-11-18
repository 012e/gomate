package middlewares

import (
	"net/http"
	"strconv"

	"github.com/012e/gomate/controllers"
	"github.com/012e/gomate/models"
	"github.com/012e/gomate/utils/json"
	"github.com/gin-gonic/gin"
)

// bind neccessary infos to the controller
func BindDefaultControllerContexts(c *controllers.DefaultController) gin.HandlerFunc {
	return func(g *gin.Context) {
		var username string = g.MustGet("username").(string)
		var user models.User
		err := c.DB.First(&user, "username = $1", username).Error
		if err != nil {
			g.AbortWithStatusJSON(http.StatusInternalServerError, json.FailUnknown())
			return
		}
		// save user infos to the controller
		userContext := &c.PermManager.Context
		userContext.Username = username
		userContext.UserID = user.ID
		userContext.UserIDStr = strconv.FormatInt(user.ID, 10)
		if user.HaveGroup {
			userContext.GroupID = user.GroupID
			userContext.GroupIDStr = strconv.FormatInt(user.GroupID, 10)
			userContext.UserHaveGroup = true
		} else {
			userContext.UserHaveGroup = false
		}
		userContext.Route = g.Request.URL.Path
	}

}
