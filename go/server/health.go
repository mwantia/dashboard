package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupHealth(r *gin.Engine) {
	r.GET("/health", RouteHealth)
}

func RouteHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}
