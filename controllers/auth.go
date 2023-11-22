package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/012e/gomate/controllers/permmanager"
	"github.com/012e/gomate/models"
	"github.com/012e/gomate/utils/resp"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type registerForm struct {
	Name     string `form:"name"     binding:"required"`
	Username string `form:"username" binding:"required,ascii"`
	Password string `form:"password" binding:"required,ascii"`
}

func isDigit(s string) bool {
	for c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func validUsername(s string) (bool, string) {
	// validator already checked for ascii only string

	if strings.Contains(s, " ") {
		return false, "username can't contain space"
	}

	if isDigit(s) {
		return false, "username can't be a number"
	}

	return true, ""
}

func userExists(db *gorm.DB, username string) (bool, error) {
	var exists bool
	err := db.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = $1", username).
		Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil

}

func addDefaultPerms(p *permmanager.PermManager, username string) error {
	err := p.AddPoliciesForUser("/todo/new", "POST")
	if err != nil {
		return err
	}

	err = p.AddPoliciesForUser("/group/new", "POST")
	if err != nil {
		return err
	}

	return nil
}

func (c DefaultController) Register(g *gin.Context) {
	var form registerForm
	if err := g.ShouldBind(&form); err != nil {
		g.JSON(http.StatusBadRequest, resp.Fail(err.Error()))
		return
	}

	logrus.Info("register user " + form.Username)
	exists, err := userExists(c.DB, form.Username)
	if exists {
		g.JSON(
			http.StatusConflict,
			resp.Fail("user already exists"),
		)
		return
	}
	logrus.Debug("checked existence of user")

	if ok, why := validUsername(form.Username); !ok {
		g.JSON(http.StatusBadRequest, resp.Fail(why))
		return
	}
	logrus.Debug("checked username")

	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(form.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		g.JSON(
			http.StatusInternalServerError,
			resp.FailUnknown(),
		)
		return
	}

	err = c.DB.Model(&models.User{}).
		Create(map[string]any{"username": form.Username, "password_hash": hashed, "name": form.Name}).Error
	if err != nil {
		g.JSON(http.StatusInternalServerError, resp.Fail("something went wrong while creating new user"))
		return
	}
	g.JSON(http.StatusOK, resp.Ok("created new user"))
	logrus.Debug("created new user")

	err = addDefaultPerms(c.PermManager, form.Username)
	if err != nil {
		g.JSON(http.StatusInternalServerError, resp.Fail("something went wrong while adding permissions for user"))
		return
	}
	logrus.Debug("added default permissions for user")
}

type loginForm struct {
	Username string `form:"username" binding:"required,ascii"`
	Password string `form:"password" binding:"required,ascii"`
}

func (c DefaultController) Login(g *gin.Context) {
	var form loginForm
	if err := g.ShouldBind(&form); err != nil {
		g.JSON(http.StatusBadRequest, resp.Fail(err.Error()))
		return
	}
	strings.TrimSpace(form.Username)

	// the response message if username or password is wrong, prevent username checking.
	const wrongInfo = "wrong username or password"

	// TODO: fix timimng attack on usernames
	exists, err := userExists(c.DB, form.Username)
	if err != nil {
		logrus.Debugf("failed with %s", err)
		g.JSON(http.StatusInternalServerError, resp.FailUnknown())
		return
	}
	if !exists {
		logrus.Debugf("failed with %s", err)
		g.JSON(http.StatusConflict, resp.Fail(wrongInfo))
		return
	}
	logrus.Debug("checked user existence")

	var user models.User
	err = c.DB.First(&user, "username = $1", form.Username).Error
	if err != nil {
		g.JSON(http.StatusInternalServerError, resp.FailUnknown())
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(form.Password)); err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError, resp.Fail(wrongInfo))
		return
	}
	logrus.Debug("finished hashing password")

	// generate new token
	token := uuid.NewString()
	expiry := time.Now().AddDate(0, 0, 7)
	err = c.DB.Model(&models.Session{}).Create(map[string]any{
		"token":    token,
		"expiry":   expiry,
		"username": user.Username,
	}).Error
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError, resp.FailUnknown())
		return
	}
	logrus.Debug("generated new session token")

	// no need to remove old token
	g.SetCookie("sessionToken", token, 99999, "/", "localhost:8080", true, true)
	g.JSON(http.StatusOK, resp.Ok("logged you in"))
}
