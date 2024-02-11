package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type fileRoute struct {
	route    string
	filepath string
}

func Setup() {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	/** Load pages & templates */
	router.LoadHTMLGlob("public/**/*.html")

	var fileRoutes = []fileRoute{
		/** Assets */
		{"/favicon.ico", "./public/assets/favicon.svg"},
		{"/assets/circle-check.svg", "./public/assets/circle-check.svg"},

		/** Javascript */
		{"/js/htmx.js", "./public/js/htmx.min.js"},

		/** Stylesheets */
		{"/css/main.css", "./public/css/main.css"},
	}

	for _, route := range fileRoutes {
		router.StaticFile(route.route, route.filepath)
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Run(":8080")

}
