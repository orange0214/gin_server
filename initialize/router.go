package initialize

import (
	"net/http"
	"project/global"
	"project/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()

	// 系统路由
	systemRouter := router.RouterGroupApp.System

	// 公开路由
	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 状态检查
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}

	return Router
}
