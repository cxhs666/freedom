package router

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Register(gr *gin.Engine) {
	// gin.Context，封装了request和response
	r = gr
	//注册具体的路由
	registerWebRouters()
	registerApiRouters()
	registerBackendRouters()
}
