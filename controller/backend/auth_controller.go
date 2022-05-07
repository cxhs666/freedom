package backend

import (
	"fmt"
	"freedom/global"
	"freedom/model"
	"freedom/service/common"
	"freedom/service/response"

	"github.com/gin-gonic/gin"
)

type LoginAccount struct {
	Account  string `form:"account" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type AuthController struct{}

var LoginErrorMsg string = "账号或密码不存在"

func (a *AuthController) Login(c *gin.Context) {
	var form LoginAccount
	if err := c.ShouldBind(&form); err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	var Admin model.Admin
	err := global.DB.Where("email = ?", form.Account).First(&Admin).Error
	if err != nil {
		response.FailWithMsg(c, LoginErrorMsg)
		return
	}

	if Admin.Password != common.EncryptByMd5(Admin.Email+form.Password) {
		response.FailWithMsg(c, LoginErrorMsg)
		return
	}

	auth := common.Auth{
		Id:   int64(Admin.Id),
		Name: Admin.Name,
	}

	res, err := common.Jwt.Create(auth)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	response.SuccessWithData(c, gin.H{"token": res, "user": gin.H{"name": Admin.Name, "avatar": Admin.Avatar}})
}

func (a *AuthController) Logout(c *gin.Context) {
	// token := c.DefaultQuery("token", "枯藤")
	// res, _ := common.Jwt.Verify(token)
	res, _ := c.Get("Auth")
	response.SuccessWithMsg(c, fmt.Sprintf("token is %+v", res))
}
