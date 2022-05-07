package bootstrap

import (
	"fmt"
	"freedom/config"
	"freedom/global"
	"freedom/model"
	"freedom/service/common"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

//配置默认超级管理员的账号信息
const defaultAdminName = "admin"
const defaultAdminEmail = "564210220@qq.com"
const defaultAdminPwd = "123456"

func InitTable(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Admin{},
	)
	if err != nil {
		os.Exit(0)
	}

	//插入超级管理
	var count int64
	db.Table("admin").Count(&count)
	if count == 0 {
		var admin model.Admin
		admin.Name = defaultAdminName
		admin.Email = defaultAdminEmail
		admin.Password = common.EncryptByMd5(admin.Email + defaultAdminPwd)

		err = db.Create(&admin).Error
		if err != nil {
			fmt.Println("admin 初始化数据失败")
			os.Exit(0)
		}
	}
}

func init() {
	//设置env环境配置
	err := godotenv.Load()
	if err != nil {
		fmt.Println("环境变量设置失败", err.Error())
	}

	//初始化数据库
	global.DB = config.DbInit()

}
