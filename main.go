package main

import (
	"os"

	"github.com/012e/gomate/controllers"
	"github.com/012e/gomate/models"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var enforcer *casbin.Enforcer
var defaultController *controllers.DefaultController

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	dsn := os.Getenv("GOMATE_DSN")
	if dsn == "" {
		panic("Can't find dsn (via GOMATE_DSN env var)")
	}
	var err error
	db, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Session{},

		&models.Group{},
		&models.GroupUser{},
		&models.GroupJoinCode{},

		&models.Todo{},
	)

	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}
	enforcer, err = casbin.NewEnforcer("rbac_model.conf", a)
	if err != nil {
		panic(err)
	}
	enforcer.LoadPolicy()
	defaultController = controllers.Setup(db, enforcer)
}

func main() {
	r := gin.Default()
	createRoutes(r)
	r.Run()
}
