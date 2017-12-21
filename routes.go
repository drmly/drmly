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
	if isJob {
		jobs := "one"
		log.WithFields(log.Fields{
			"job": "mp42dream",
		}).Info("added new job to queue")
		c.HTML(200, "jobs.html", gin.H{
			jobs: jobs,
		})
	} else {
		dream(c)
		c.HTML(200, "jobs.html", gin.H{})
	}
}
func getAbout(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"variableName": "value",
	})
}
func getCode(c *gin.Context) {
	c.HTML(200, "code.html", gin.H{
		"variableName": "value",
	})
}

func getJobs(c *gin.Context) {
	c.HTML(200, "jobs.html", gin.H{
		"jobs": "the job you just made is awesome!",
		"est":  "should be done in ten zetaseconds",
	})
}
