package main

import (
	"math/rand"
	"net/http"	
	"time"
	
	"github.com/skratchdot/open-golang/open"
	"github.com/gin-gonic/gin"
	"github.com/glibs/gin-webserver"
)

func web() {
	host := "0.0.0.0:8080"
	server := InitializeServer(host)
	server.Start()
	log.Info("Gin web server started on " + host)
}

// InitializeServer gets our gin running front end poppin off
func InitializeServer(host string) (server *network.WebServer) {
	rand.Seed(time.Now().UTC().UnixNano())
	// Make sure folders exist that we want:
	if err := ensureBindDirs(); err != nil {
		log.Error("Failed to have home working dir to put the files into at ~/Desktop/bind, suxx", err)
	} else {
		log.Info("bind dirs ensured!")
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
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


	return network.InitializeWebServer(r, host)
}
