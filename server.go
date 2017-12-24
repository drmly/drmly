package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/glibs/gin-webserver"
	"github.com/olahol/melody"
	"github.com/skratchdot/open-golang/open"
)

func web() {
	host := "0.0.0.0:8080"
	server := InitializeServer(host)
	server.Start()
	Log.Info("Gin web server started on " + host)
}

// InitializeServer starts our gin server
func InitializeServer(host string) (server *network.WebServer) {
	// Make sure folders exist that we want:
	if err := ensureBindDirs(); err != nil {
		Log.Error("Failed to have home working dir to put the files into at ~/Desktop/bind, err: ", err)
	} else {
		Log.Info("bind dirs ensured!")
	}
	if os.Args[0] != "d" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	mel := melody.New() // melody middleware

	// root route ('/')
	r.GET("/", func(ctx *gin.Context) {
		// serve a static HTML file
		http.ServeFile(ctx.Writer, ctx.Request, "public/tmpl/index.html")
	})

	// websocket route
	r.GET("/ws", func(ctx *gin.Context) {
		// handle request with Melody
		mel.HandleRequest(ctx.Writer, ctx.Request)
	})

	// Melody message handler
	mel.HandleMessage(func(ses *melody.Session, msg []byte) {
		// broadcast message to connected sockets
		mel.Broadcast(msg)
	})

	r.LoadHTMLGlob("public/tmpl/*.html")
	r.StaticFS("/videos", http.Dir(basePath+"/videos"))
	r.StaticFS("/frames", http.Dir(basePath+"/frames"))
	r.Static("/public", "./public")

	r.POST("/g", postIndex)
	r.GET("/about", getAbout)
	r.GET("/jobs", getJobs)
	r.GET("/code", getCode)
	r.GET("/openframes", func(c *gin.Context) {
		open.Run(basePath + "/frames")
	})
	r.GET("/openvideos", func(c *gin.Context) {
		open.Run(basePath + "/videos")
	})
	r.GET("/openlogs", func(c *gin.Context) {
		open.Run(basePath + "/logs")
	})
	r.GET("/toggleClipYt", func(c *gin.Context) {
		open.Run(basePath + "/logs")
	})
	// start the worker
	go worker(jobChan)

	return network.InitializeWebServer(r, host)
}
