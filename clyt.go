package main

import (
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/tjgq/clipboard"
)

var ytURL string //we use a seperate chan so we can filter out and ignore bad URL's ytdl can't download

func init() {
	ytChan := make(chan string, 1)
	go func() {
		for {
			clipboard.Notify(ytChan)
			got := <-ytChan
			cmd, err := exec.Command("youtube-dl", "-s", got).CombinedOutput()
			if err != nil {
				Log.WithFields(logrus.Fields{
					"event": "ytdl",
					"error": err,
				}).Error("that was just not a good url from clipboad")
				continue
			} else {
				Log.Info("youtube-dl said: ", string(cmd))
				fmt.Println("got : ", got)
				ytURL =  got
			}
		}
	}()
}
