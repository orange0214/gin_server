package system

import (
	"fmt"
	systemReq "project/system/request"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	// err := c.ShouldBindJSON(&l) // 将请求中的Json赋给结构体l
	// key := c.ClientIP()         // 获得请求头部信息

	// 处理错误
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	fmt.Print(l)
}
