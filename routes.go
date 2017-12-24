package main

import (
	"sync"

	"github.com/gin-gonic/gin"
)

func getIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"v": version,
	})
}

func worker(jobChan <-chan *gin.Context) {
    defer wg.Done()
    for job := range jobChan {
		Dream(job)
		Log.Info("finished a job")
    }
}

// make a channel with a capacity of 100.
var jobChan chan *gin.Context

func init()  {
	jobChan = make(chan *gin.Context, 8)
}
var wg sync.WaitGroup

func postIndex(c *gin.Context) {
	Log.Info("Is this exist?")
	cCp:=c.Copy()
	wg.Add(1)
	jobChan <- cCp
	wg.Wait()
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
