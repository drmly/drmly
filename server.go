package main

import (
	"math/rand"

	log "github.com/sirupsen/logrus"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/glibs/gin-webserver"
)

func web() {
	host := "0.0.0.0:8080"
	server := InitializeServer(host)
	server.Start()
	log.Println("Gin web server started on " + host)
	for {
		time.Sleep(10 * time.Minute)
		// Hold program open to test web server
	}
	// After the server is not useful
	// server.Stop()
	// server.Listener.TCPListener.

}

// InitializeServer gets our gin running front end poppin off
func InitializeServer(host string) (server *network.WebServer) {
	rand.Seed(time.Now().UTC().UnixNano())
	newLogger()
	// Make sure folders exist that we want:

	if err := ensureDreamlyDirs(); err != nil {
		log.Error("Failed to have home working dir to put the files into at ~/Desktop/dreamly, suxx", err)
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r.LoadHTMLBinData(AssetNames(), MustAsset)
	r.LoadHTMLGlob("frontend/templates/*.html")
	r.Static("/public/css/", "./public/css")
	r.Static("/public/js/", "./public/js/")
	r.Static("/public/fonts/", "./public/fonts/")
	r.Static("/public/img/", "./public/img/")
	// usr, err := homedir.Dir()
	// if err != nil {
	// 	log.Error("failed to get homedir", err)
	// }
	// exp, err := homedir.Expand(usr)
	// if err != nil {
	// 	log.Error("failed to get expanded homedir", err)
	// }
	// // log.Info("frames dir path for dreamly app is ", exp+"/frames/")
	// path := fmt.Sprintf("%s/Desktop/frames", exp)
	// log.Info("")
	r.StaticFile("favicon.ico", "./favicon.ico")

	r.GET("/", getIndex)
	r.POST("/g", postIndex)
	r.GET("/g", getAbout)
	r.GET("/downloads", getDownloads)
	r.GET("/about", getAbout)
	r.GET("/contact", getContact)
	r.GET("/jobs", getJobs)
	r.GET("/code", getCode)
	r.GET("/donate", getDonate)

	return network.InitializeWebServer(r, host)
}
