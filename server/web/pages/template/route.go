package template

import (
	"github.com/gin-gonic/gin"
)

func RouteWebPages(route *gin.RouterGroup) {
	route.GET("/", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", Indexhtml)
	})

	route.GET("/android-chrome-192x192.png", func(c *gin.Context) {
		c.Data(200, "image/png", Androidchrome192x192png)
	})

	route.GET("/android-chrome-512x512.png", func(c *gin.Context) {
		c.Data(200, "image/png", Androidchrome512x512png)
	})

	route.GET("/apple-touch-icon.png", func(c *gin.Context) {
		c.Data(200, "image/png", Appletouchiconpng)
	})

	route.GET("/asset-manifest.json", func(c *gin.Context) {
		c.Data(200, "application/json", Assetmanifestjson)
	})

	route.GET("/browserconfig.xml", func(c *gin.Context) {
		c.Data(200, "text/xml; charset=utf-8", Browserconfigxml)
	})

	route.GET("/dlnaicon-120.jpg", func(c *gin.Context) {
		c.Data(200, "image/jpeg", Dlnaicon120jpg)
	})

	route.GET("/dlnaicon-120.png", func(c *gin.Context) {
		c.Data(200, "image/png", Dlnaicon120png)
	})

	route.GET("/dlnaicon-48.jpg", func(c *gin.Context) {
		c.Data(200, "image/jpeg", Dlnaicon48jpg)
	})

	route.GET("/dlnaicon-48.png", func(c *gin.Context) {
		c.Data(200, "image/png", Dlnaicon48png)
	})

	route.GET("/favicon-16x16.png", func(c *gin.Context) {
		c.Data(200, "image/png", Favicon16x16png)
	})

	route.GET("/favicon-32x32.png", func(c *gin.Context) {
		c.Data(200, "image/png", Favicon32x32png)
	})

	route.GET("/favicon.ico", func(c *gin.Context) {
		c.Data(200, "image/vnd.microsoft.icon", Faviconico)
	})

	route.GET("/index.html", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", Indexhtml)
	})

	route.GET("/mstile-150x150.png", func(c *gin.Context) {
		c.Data(200, "image/png", Mstile150x150png)
	})

	route.GET("/site.webmanifest", func(c *gin.Context) {
		c.Data(200, "application/manifest+json", Sitewebmanifest)
	})

	route.GET("/static/js/2.a6a6e0da.chunk.js", func(c *gin.Context) {
		c.Data(200, "text/javascript; charset=utf-8", Staticjs2a6a6e0dachunkjs)
	})

	route.GET("/static/js/2.a6a6e0da.chunk.js.LICENSE.txt", func(c *gin.Context) {
		c.Data(200, "text/plain; charset=utf-8", Staticjs2a6a6e0dachunkjsLICENSEtxt)
	})

	route.GET("/static/js/2.a6a6e0da.chunk.js.map", func(c *gin.Context) {
		c.Data(200, "application/json", Staticjs2a6a6e0dachunkjsmap)
	})

	route.GET("/static/js/main.6dafc43b.chunk.js", func(c *gin.Context) {
		c.Data(200, "text/javascript; charset=utf-8", Staticjsmain6dafc43bchunkjs)
	})

	route.GET("/static/js/main.6dafc43b.chunk.js.map", func(c *gin.Context) {
		c.Data(200, "application/json", Staticjsmain6dafc43bchunkjsmap)
	})

	route.GET("/static/js/runtime-main.33603a80.js", func(c *gin.Context) {
		c.Data(200, "text/javascript; charset=utf-8", Staticjsruntimemain33603a80js)
	})

	route.GET("/static/js/runtime-main.33603a80.js.map", func(c *gin.Context) {
		c.Data(200, "application/json", Staticjsruntimemain33603a80jsmap)
	})
}
