package controllers

import (
	"net/http"
	"strconv"

	"github.com/012e/gomate/models"
	"github.com/012e/gomate/utils/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (c *DefaultController) JoinGroup(g *gin.Context) {
	// groupJoinCode := g.Param("code")
}

func (c *DefaultController) LeaveGroup(g *gin.Context) {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		err := c.DB.
			Model(&models.User{}).
			Where(c.PermManager.Context.UserID).
			Updates(map[string]any{"have_group": false, "group_id": nil}).
			Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		g.JSON(http.StatusInternalServerError, resp.FailUnknow())
		return
	}
	g.JSON(http.StatusAccepted, resp.Ok("left your group"))

}

func (c *DefaultController) DeleteGroup(g *gin.Context) {}

type createGroupForm struct {
	Name        string `form:"name" binding:"omitempty,required"`
	Description string `form:"description"`
}

func (c *DefaultController) CreateGroup(g *gin.Context) {
	var form createGroupForm
	if err := g.ShouldBind(&form); err != nil {
		g.JSON(http.StatusBadRequest, resp.Fail(err.Error()))
		return
	}

	err := c.DB.Transaction(func(tx *gorm.DB) error {
		tx = tx.Session(&gorm.Session{})
		// create group
		var group models.Group
		group.Name = form.Name
		group.Description = form.Description
		err := tx.Select("name", "description").Create(&group).Error
		if err != nil {
			return err
		}
		logrus.Debugf("created new group with id %d", group.ID)

		// user join group
		err = tx.Model(&models.GroupUser{}).Create(map[string]any{
			"group_id": group.ID,
			"user_id":  c.PermManager.Context.UserID,
		}).Error
		if err != nil {
			return err
		}
		logrus.Debug("user joined new group")

		// user derives permissions from group
		groupIDStr := strconv.FormatInt(group.ID, 10)
		err = c.PermManager.UserDerivePolicesFromGroup(groupIDStr)
		if err != nil {
			return err
		}
		logrus.Debug("user got roles of new group")

		// update user info
		err = tx.
			Model(&models.User{}).
			Where(c.PermManager.Context.UserID).
			Updates(map[string]any{
				"have_group": true,
				"group_id":   group.ID,
			}).Error
		if err != nil {
			return err
		}

		// transaction is successful
		return nil
	})

	if err != nil {
		logrus.Debugf("failed to create group")
		g.JSON(http.StatusInternalServerError, resp.FailUnknow())
		return
	}
	g.JSON(http.StatusAccepted, resp.Ok("created new group"))
}

func (c *DefaultController) CreateJoinCode(g *gin.Context) {
}
