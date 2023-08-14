package system

import (
	"project/global"
	"project/model/common/response"
	system "project/model/system"
	systemReq "project/model/system/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct{}

// Login
// @Router	/base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l) // 将请求中的Json赋给结构体l
	// key := c.ClientIP()         // 获得请求头部信息

	// 处理错误
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	u := &system.SysUser{Username: l.Username, Password: l.Password}
	user, err := userService.Login(u)
	if err != nil {
		global.GVA_LOG.Error("登录失败！用户名不存在或密码错误！", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	if user.Enable != 1 {
		global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
		response.FailWithMessage("用户被禁止登录", c)
		return
	}
	// b.TokenNext(c, *user)

	return
}
