package middlerware

import (
	"fmt"
	"freedom/global"
	"freedom/model"
	"freedom/service/common"

	"github.com/gin-gonic/gin"
)

//全局中间件
type Global struct{}

//局部中间件
type Part struct{}

type Middlerware struct {
	Global
	Part
}

var Md = new(Middlerware)

func (p *Part) AdminLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")

		if len(tokenStr) == 0 {
			unauthorizedRep(c)
			return
		}
		var token string
		n, _ := fmt.Sscanf(tokenStr, "Bearer %s", &token)

		if n == 0 {
			unauthorizedRep(c)
			return
		}

		loginAuth, err := common.Jwt.Verify(token)
		if err != nil {
			unauthorizedRep(c)
			return
		}

		var Admin model.Admin
		err = global.DB.First(&Admin, loginAuth.Id).Error

		if err != nil {
			unauthorizedRep(c)
			return
		}

		//当前登录用户信息
		auth := Admin
		c.Set("Auth", auth)

		c.Next()
	}
}

func unauthorizedRep(c *gin.Context) {
	c.JSON(401, "请登录")
	c.Abort()
}
