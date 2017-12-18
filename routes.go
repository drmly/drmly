package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"v": version,
	})
}
func postIndex(c *gin.Context) {
	log.Info("Is this exist?")
	mp4ToDream(c)
	getIndex(c)
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
	c.HTML(200, "downloads.html", gin.H{
		"variableName": "value",
	})
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
