package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	initializeRoutes()

	// Start serving the application
	_ = router.Run()

}

// Rendering of HTML,JSON or XML based on "Accept" header
// if header does not specify it will render html

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with Json
		c.JSON(http.StatusOK, data["payload"])
	case "application/XML":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)

	}
}
