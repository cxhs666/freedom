package router

import (
	bc "freedom/controller"
)

const apiGroup = "/api"

func registerApiRouters() {
	controller := bc.ControllerHandler.ApiController
	api := r.Group(apiGroup)
	{
		api.GET("/user", controller.UserController.User)
	}
}
