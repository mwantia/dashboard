package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupHealth(r *gin.Engine) error {
	r.GET("/health", RouteHealth)
	return nil
}

func RouteHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}
