package server

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/mwantia/dashboard/templates/layout"
)

func SetupHttp(r *gin.Engine) error {
	r.Static("/static", "./src/static")

	r.GET("/home", RouteHome)
	r.NoRoute(RouteHome)

	return nil
}

func RouteHome(ctx *gin.Context) {
	NavigateView(ctx, layout.Main())
}

func NavigateView(ctx *gin.Context, component templ.Component) {
	ctx.Status(http.StatusOK)
	if err := component.Render(ctx.Request.Context(), ctx.Writer); err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
}

func HtmxRequest(ctx *gin.Context) bool {
	header := ctx.GetHeader("HX-Request")
	return len(header) > 0 && header == "true"
}
