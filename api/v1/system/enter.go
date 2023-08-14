package system

import "project/service"

type ApiGroup struct {
	BaseApi // 用户登录
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
