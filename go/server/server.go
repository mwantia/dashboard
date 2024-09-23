package server

import "github.com/gin-gonic/gin"

func Start() {
	r := SetupRouter()
	// Health Setup
	SetupHealth(r)
	// Start listening and serving requests
	if err := r.Run(":9000"); err != nil {
		panic(err)
	}
}

func SetupRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	return router
}
