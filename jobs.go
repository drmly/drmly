package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var maxJobs int
var jobLog = logrus.New()

func newJob(name string) {
	file, err := os.OpenFile(basePath+"/logs/"+name+".log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		jobLog.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	jobLog.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	  }).Info("A group of walrus emerges from the ocean")
	jobLog.Info("hey we have job logs now, cool!")
}

// let's add a queue and job cancellation
//  and let's save job data in files in /logs that way people can keep track of jobs done and see what happened during the job (e.g. randomized frame every tenth frame)
