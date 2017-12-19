package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
)

func getIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"v": version,
	})
}
func postIndex(c *gin.Context) {
	log.Info("Is this exist?")
	if isJob {
		log.WithFields(log.Fields{
			"job": "mp42dream",
		}).Info("was denied a new job, job already running")
		c.String(200, "a job already running, wait til it's finished") //todo implement a queue here instead
		return
	} else {
		dream(c)
	}
	// Design the flow and run it
	// flow := run.Sequence(
	// 	run.Parallel(),

	// )

	// ctx := floc.NewContext()

	// ctrl := floc.NewControl(ctx)

	// floc.RunWith(ctx, ctrl, flow)

}
func getAbout(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"variableName": "value",
	})
}

func getContact(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{
		"variableName": "value",
	})
}

func getDownloads(c *gin.Context) {
	open.Run(basePath)
	open.Run(basePath + "/videos")
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
func getJobs(c *gin.Context) {
	c.HTML(200, "jobs.html", gin.H{
		"variableName": "value",
	})
}
func getCode(c *gin.Context) {
	c.HTML(200, "code.html", gin.H{
		"variableName": "value",
	})
}
func getDonate(c *gin.Context) {
	c.HTML(200, "donate.html", gin.H{
		"variableName": "value",
	})
}
