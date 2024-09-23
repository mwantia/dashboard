package server

import "github.com/gin-gonic/gin"

func Start() {
	r := SetupRouter()
	// Health Setup
	if err := SetupHealth(r); err != nil {
		panic(err)
	}
	// API Setup
	if err := SetupApi(r); err != nil {
		panic(err)
	}
	// HTTP Setup
	if err := SetupHttp(r); err != nil {
		panic(err)
	}
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
