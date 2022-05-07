package backend

import (
	"errors"
	"fmt"
	"freedom/global"
	"freedom/model"
	"freedom/service/common"
	"freedom/service/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminController struct{}

type AdminForm struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func (u *AdminController) User(c *gin.Context) {
	id := c.Param("id")

	var admin model.Admin
	err := global.DB.Where("id = ?", id).First(&admin).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMsg(c, "用户不存在")
		return
	}

	response.SuccessWithData(c, admin)
}

func (u *AdminController) Users(c *gin.Context) {
	var users model.Admin

	if res := global.DB.Find(&users); res.Error != nil {
		response.Fail(c)
		return
	}

	response.SuccessWithData(c, users)
}

func (u *AdminController) Create(c *gin.Context) {
	var form AdminForm
	if err := c.ShouldBind(&form); err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	var admin model.Admin
	err := global.DB.Where("email = ?", form.Email).First(&admin).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMsg(c, "新增失败")
		return
	}
	if admin.Email != "" {
		response.FailWithMsg(c, "当前邮箱已经存在")
		return
	}

	admin.Email = form.Email
	admin.Name = form.Name
	admin.Password = common.EncryptByMd5(form.Email + form.Password)

	err = global.DB.Create(&admin).Error
	if err != nil {
		fmt.Println(err.Error())
		response.FailWithMsg(c, "新增失败")
		return
	}
	response.SuccessWithMsg(c, "新增成功")
}
