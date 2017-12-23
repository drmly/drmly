package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/glibs/gin-webserver"
	"github.com/skratchdot/open-golang/open"
)

func web() {
	host := "0.0.0.0:8080"
	server := InitializeServer(host)
	server.Start()
	Log.Info("Gin web server started on " + host)
}

// InitializeServer gets our gin running front end poppin off
func InitializeServer(host string) (server *network.WebServer) {
	rand.Seed(time.Now().UTC().UnixNano())
	// Make sure folders exist that we want:
	if err := ensureBindDirs(); err != nil {
		Log.Error("Failed to have home working dir to put the files into at ~/Desktop/bind, err: ", err)
	} else {
		Log.Info("bind dirs ensured!")
	}
	if os.Args[0] != "d" { //development mode
		gin.SetMode(gin.ReleaseMode)
	} 
	r := gin.New()
	r.LoadHTMLGlob("public/tmpl/*.html")
	r.StaticFS("/videos", http.Dir(basePath+"/videos"))
	r.StaticFS("/frames", http.Dir(basePath+"/frames"))
	r.Static("/public", "./public")
	r.GET("/", getIndex)
	r.POST("/g", postIndex)
	r.GET("/g", getIndex)
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

	return network.InitializeWebServer(r, host)
}
