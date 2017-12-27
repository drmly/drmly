package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var maxJobs int
var jobLog = logrus.New()

func newJobLog(name string) {
	file, err := os.OpenFile(basePath+"/logs/"+name+".txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		jobLog.Out = file
	} else {
		Log.Info("Failed to log to file, using default stderr")
		jobLog.Out = os.Stderr
	}
}

// let's add a queue and job cancellation
//  and let's save job data in files in /logs that way people can keep track of jobs done and see what happened during the job (e.g. randomized frame every tenth frame)
