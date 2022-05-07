package router

import (
	bc "freedom/controller"
	"freedom/middlerware"
)

const backendGroup = "/sys"

func registerBackendRouters() {
	controller := bc.ControllerHandler.BackController
	md := middlerware.Md
	api := r.Group(backendGroup)
	{
		api.POST("/login", controller.AuthController.Login)
	}

	//验证登录中间件
	api.Use(md.AdminLogin())
	{

		api.GET("/logout", controller.AuthController.Logout)

		api.GET("/user/:id", controller.AdminController.User)
		api.GET("/users", controller.AdminController.Users)
		api.POST("/users/create", controller.AdminController.Create)
	}

}
