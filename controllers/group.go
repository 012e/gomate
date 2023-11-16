package controllers

import "github.com/gin-gonic/gin"

func (c *DefaultController) JoinGroup(g *gin.Context) {
	// groupJoinCode := g.Param("code")
}

func (c *DefaultController) LeaveGroup(g *gin.Context) {
}

func (c *DefaultController) DeleteGroup(g *gin.Context) {}

type createGroupForm struct {
	Name        string `form:"name" binding:"omitempty,required"`
	Description string `form:"description"`
}
func (c *DefaultController) CreateGroup(g *gin.Context) {

}

func (c *DefaultController) CreateJoinCode(g *gin.Context) {
}

