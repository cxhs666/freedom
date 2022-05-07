package api

import "github.com/gin-gonic/gin"

type UserController struct{}

func (u *UserController) User(c *gin.Context) {
	c.String(200, "hi hello world api")
}
