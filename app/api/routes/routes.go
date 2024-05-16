package routes

import (
	"gosplitwise/app/api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {
	o := g.Group("/o")
	// r := g.Group("/r")
	o.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	o.POST("/user/register", handlers.Register)
	o.POST("/user/login", handlers.Login)
}
