package controller

import (
	"freedom/controller/api"
	"freedom/controller/backend"
)

type Controller struct {
	BackController backend.BackendController
	ApiController  api.ApiController
}

var ControllerHandler = new(Controller)
