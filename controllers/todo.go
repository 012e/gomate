package controllers

import (
	"fmt"
	"net/http"

	"github.com/012e/gomate/models"
	"github.com/012e/gomate/utils/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// TODO: If finished blank still work. It should response an invalid request instead.
type todoForm struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	Finished    bool   `form:"finished" binding:"omitempty,boolean"`
}

func (c *DefaultController) CreateTodo(ctx *gin.Context) {
	logrus.Info("creating new todo")
	var todoForm todoForm
	if err := ctx.ShouldBind(&todoForm); err != nil {
		ctx.JSON(http.StatusBadRequest, json.Fail("invalid create todo form: "+err.Error()))
		return
	}

	var todo models.Todo
	err := c.DB.Model(&todo).Create(map[string]any{
		"title":       todoForm.Title,
		"description": todoForm.Description,
		"finished":    todoForm.Finished,
	}).Error
	if err != nil {
		ctx.JSON(http.StatusTeapot, json.Fail("something went wrong: "+err.Error()))
		return
	}
	c.PermManager.AddPoliciesForUser(fmt.Sprintf("/todo/%d", todo.ID), "GET", "PATCH", "DELETE")
	ctx.JSON(http.StatusOK, json.Ok("created new todo"))

	logrus.Info("created new todo")
}

// func (c DefaultController) DeleteTodo(ctx *gin.Context) {
// 	todoID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong parsing todo id"})
// 		return
// 	}
// 	logrus.Infof("deleting todo id %d", todoID)
// 	err = models.DeleteTodoByID(c.DB, todoID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "deleting todo with id " + ctx.Param("id")})
// 		return
// 	}
// 	ctx.JSON(http.StatusAccepted, gin.H{"success": "deleted todo with id " + ctx.Param("id")})
// }
