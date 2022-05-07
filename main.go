package main

import (
	"freedom/global"
	"freedom/router"
	"time"

	"freedom/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {

	if global.DB != nil {

		//初始化数据表
		bootstrap.InitTable(global.DB)
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()

		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		db.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		db.SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		db.SetConnMaxLifetime(time.Hour)
	}

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	router.Register(r)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
