package response

import "github.com/gin-gonic/gin"

const (
	ERROR   = 1
	SUCCESS = 0
)

type List struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type ResponseData struct {
	ErrorCode int         `json:"errorCode"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

func result(c *gin.Context, code int, errorCode int, msg string, data interface{}) {
	c.JSON(code, ResponseData{errorCode, msg, data})
}

func Success(c *gin.Context) {
	result(c, 200, SUCCESS, "操作成功", "")
}

func SuccessWithMsg(c *gin.Context, msg string) {
	result(c, 200, SUCCESS, msg, "")
}

func SuccessWithData(c *gin.Context, data interface{}) {
	result(c, 200, SUCCESS, "ok", data)
}

func Fail(c *gin.Context) {
	result(c, 200, ERROR, "操作失败", "")
}

func FailWithMsg(c *gin.Context, msg string) {
	result(c, 200, ERROR, msg, "")
}

func FailWithCode(c *gin.Context, errorCode int, msg string) {
	result(c, 200, errorCode, msg, "")
}

func FailWithData(c *gin.Context, data interface{}) {
	result(c, 200, ERROR, "fail", data)
}
