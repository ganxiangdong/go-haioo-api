package controllers

import (
	"github.com/gin-gonic/gin"
	"haioo/code"
)

func init() {
}

//声明控制器基类
type BaseController struct {
}

//控制器基类的方法
func (this BaseController) Show404(c *gin.Context) {
	c.String(404, "～～～～页面没有找到，我们正在殴打程序员中")
}

//统一输出方法
func (this BaseController) ResponseData(c *gin.Context, codeItem code.CodeItem, msg string, result interface{}) {
	response := map[string]interface{}{"code": codeItem.Code, "description": codeItem.Description, "msg": msg, "result": result}
	c.JSON(200, response)
}
