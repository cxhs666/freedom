package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerWebRouters() {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!1232233")
	})
}
