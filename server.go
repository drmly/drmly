package main

import (
	"math/rand"

	log "github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/glibs/gin-webserver"
)

func web() {
	host := "0.0.0.0:8080"
	server := InitializeServer(host)
	server.Start()
	log.Println("Gin web server started on " + host)
}

// InitializeServer gets our gin running front end poppin off
func InitializeServer(host string) (server *network.WebServer) {
	rand.Seed(time.Now().UTC().UnixNano())
	newLogger()
	// Make sure folders exist that we want:
	if err := ensureDreamlyDirs(); err != nil {
		log.Error("Failed to have home working dir to put the files into at ~/Desktop/dreamly, suxx", err)
	} else {
		log.Info("dreamly dirs ensuered!")
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("public/tmpl/*.html")
	r.StaticFile("favicon.ico", "./favicon.ico")

	r.GET("/", getIndex)
	r.POST("/g", postIndex)
	r.GET("/g", getIndex)
	r.GET("/downloads", getDownloads)
	r.GET("/about", getAbout)
	r.GET("/contact", getContact)
	r.GET("/jobs", getJobs)
	r.GET("/code", getCode)
	r.GET("/donate", getDonate)
	r.GET("/frames", func(c *gin.Context) {
		open.Run(basePath + "/frames")
	})
	

	return network.InitializeWebServer(r, host)
}
