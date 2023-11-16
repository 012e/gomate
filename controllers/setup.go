package controllers

import (
	"net/http"

	"github.com/012e/gomate/controllers/permmanager"
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

func (c DefaultController) Hello(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"message": "Hello world"})
}
