package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
