package controllers

import (
	"net/http"

	"github.com/012e/gomate/controllers/permmanager"
	"github.com/012e/gomate/utils/resp"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DefaultController struct {
	*gorm.DB
	*permmanager.PermManager
}

func Setup(db *gorm.DB, e *casbin.Enforcer) *DefaultController {
	return &DefaultController{DB: db, PermManager: &permmanager.PermManager{Enforcer: e}}
}

// @Summary this is not sum
// @Schemes
// @Description do ping
// @Tags testingtag
// @Accept json
// @Produce json
// @Success 200 {object} resp.BaseOk
// @Router / [get]
func (c DefaultController) Hello(g *gin.Context) {
	g.JSON(http.StatusOK, resp.Ok("eheheh"))
}
