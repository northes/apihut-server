package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Code uint64

const (
	CodeSuccess Code = 2000
	CodeError   Code = 5000
)

var CodeMsgMap = map[Code]string{
	CodeSuccess: "成功",
	CodeError:   "失败",
}

func (c Code) Msg() string {
	if msg, ok := CodeMsgMap[c]; ok {
		return msg
	}
	return ""
}

func Success(c *gin.Context) {
	SuccessWithData(c, nil)
}

func SuccessWithData(c *gin.Context, data interface{}) {
	code := CodeSuccess
	JSON(c, code, code.Msg(), data)
}

func Error(c *gin.Context) {
	ErrorWithMsg(c, CodeError.Msg())
}

func ErrorWithMsg(c *gin.Context, msg string) {
	code := CodeError
	JSON(c, code, msg, nil)
}

func JSON(c *gin.Context, code Code, msg string, data interface{}) {
	c.JSON(http.StatusOK, &Body{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}