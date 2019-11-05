package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {

	// default router from Gin
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	api:= router.Group("/api")
	{
		api.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})

		})
	}

	_ = router.Run(":3000")
}
