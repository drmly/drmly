package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/radovskyb/watcher"
	"github.com/sirupsen/logrus"
	filetype "gopkg.in/h2non/filetype.v1"
)

type ss struct {
	path, name, it, oc, la, ch, os string
}

func (s *ss) screenDream() {
	s.it = strconv.Itoa(random(1, 40))
	s.oc = strconv.Itoa(random(1, 8))
	s.os = fmt.Sprintf("1.%s", strconv.Itoa(random(1, 9)))
	for k := range layerChannels {
		s.la = k
		s.ch = strconv.Itoa(random(1, layerChannels[k]-1))
		break
	}
	Log.Info(fmt.Sprintf("ch %s it %s oc %s la  %s os %s", s.ch, s.it, s.oc, s.la, s.os))
	Log.Info("entering dreamer goroutine")
	// deep dream the frames
	cmd, err := exec.Command("python3", "folder.py", "--input", s.path, "-ch", s.ch, "-it", s.it, "-oc", s.oc, "-la", s.la, "-os", s.os).CombinedOutput()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"event": "folder.py",
		}).Error("failed to dream", err)
		z := fmt.Sprintf("FAIL: python borked: %s", err.Error())
		mel.Broadcast([]byte(z))
	}
	Log.Info("done w/ dream loop, python said: ", string(cmd))
	savedFile := basePath + "/screens/" + s.name + "/output/1.png"
	fmt.Println(savedFile)
	//  s.path+"it"+s.it+"oc"+s.oc+"ch"+s.ch+"os"+s.os+"la"+s.la+".jpg").CombinedOutput()
	la := strings.Replace(s.la, "/", "", 1) //we want a file not to make a /conv directory or w/e
	_, err = exec.Command("ffmpeg", "-i", savedFile, s.path+"ch"+s.ch+"it"+s.it+"oc"+s.oc+"la"+la+".jpg").CombinedOutput()
	if err != nil {
		Log.Error("failed to jpg the png", err)
	}
}

func (s *ss) watch() {
	w := watcher.New()

	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(1)

	// Only notify rename and move events.
	w.FilterOps(watcher.Create)

	go func() {
		for {
			select {
			case event := <-w.Event:
				Log.Info("event path:", event.Path)
				buf, _ := ioutil.ReadFile(event.Path)
				if filetype.IsImage(buf) {
					Log.Info("we got a new screenshot file")
					s.name = haiku.Haikunate()
					s.path = fmt.Sprint(basePath + "/screens/" + s.name)
					err := os.Mkdir(s.path, 0700)
					if err != nil {
						Log.Error("couldn't make screenshot job dir", err)
						return
					}
					err = os.Mkdir(s.path+"/output", 0700)
					if err != nil {
						Log.Error("couldn't make screenshot job output dir", err)
						return
					}
					err = os.Rename(event.Path, s.path+"/"+s.name+".png")
					if err != nil {
						Log.Info("event.Path: ", event.Path)
						Log.Error("couldn't move screenshot to screens/"+s.name, err)
						return
					}
					s.screenDream()
					fmt.Println(event) // Print the event's info.
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()
	user, err := user.Current()
	if err != nil {
		Log.Fatal(err)
	}
	// Watch this folder for changes.
	if err := w.Add(user.HomeDir + "/Desktop"); err != nil {
		log.Fatalln(err)
	}
	// Start the watching process - it'll check for changes every 5000ms.
	if err := w.Start(time.Millisecond * 5000); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	s := ss{}
	go s.watch()
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

var layerChannels = map[string]int{
	"conv2d0_pre_relu/conv":                64,
	"conv2d1_pre_relu/conv":                64,
	"conv2d2_pre_relu/conv":                192,
	"mixed3a_1x1_pre_relu/conv":            64,
	"mixed3a_3x3_bottleneck_pre_relu/conv": 96,
	"mixed3a_3x3_pre_relu/conv":            128,
	"mixed3a_5x5_bottleneck_pre_relu/conv": 16,
	"mixed3a_5x5_pre_relu/conv":            32,
	"mixed3a_pool_reduce_pre_relu/conv":    32,
	"mixed3b_1x1_pre_relu/conv":            128,
	"mixed3b_3x3_bottleneck_pre_relu/conv": 128,
	"mixed3b_3x3_pre_relu/conv":            192,
	"mixed3b_5x5_bottleneck_pre_relu/conv": 32,
	"mixed3b_5x5_pre_relu/conv":            96,
	"mixed3b_pool_reduce_pre_relu/conv":    64,
	"mixed4a_1x1_pre_relu/conv":            192,
	"mixed4a_3x3_bottleneck_pre_relu/conv": 96,
	"mixed4a_3x3_pre_relu/conv":            204,
	"mixed4a_5x5_bottleneck_pre_relu/conv": 16,
	"mixed4a_5x5_pre_relu/conv":            48,
	"mixed4a_pool_reduce_pre_relu/conv":    64,
	"mixed4b_1x1_pre_relu/conv":            160,
	"mixed4b_3x3_bottleneck_pre_relu/conv": 112,
	"mixed4b_3x3_pre_relu/conv":            224,
	"mixed4b_5x5_bottleneck_pre_relu/conv": 24,
	"mixed4b_5x5_pre_relu/conv":            64,
	"mixed4b_pool_reduce_pre_relu/conv":    64,
	"mixed4c_1x1_pre_relu/conv":            128,
	"mixed4c_3x3_bottleneck_pre_relu/conv": 128,
	"mixed4c_3x3_pre_relu/conv":            256,
	"mixed4c_5x5_bottleneck_pre_relu/conv": 24,
	"mixed4c_5x5_pre_relu/conv":            64,
	"mixed4c_pool_reduce_pre_relu/conv":    64,
	"mixed4d_1x1_pre_relu/conv":            112,
	"mixed4d_3x3_bottleneck_pre_relu/conv": 144,
	"mixed4d_3x3_pre_relu/conv":            288,
	"mixed4d_5x5_bottleneck_pre_relu/conv": 32,
	"mixed4d_5x5_pre_relu/conv":            64,
	"mixed4d_pool_reduce_pre_relu/conv":    64,
	"mixed4e_1x1_pre_relu/conv":            256,
	"mixed4e_3x3_bottleneck_pre_relu/conv": 160,
	"mixed4e_3x3_pre_relu/conv":            320,
	"mixed4e_5x5_bottleneck_pre_relu/conv": 32,
	"mixed4e_5x5_pre_relu/conv":            128,
	"mixed4e_pool_reduce_pre_relu/conv":    128,
	"mixed5a_1x1_pre_relu/conv":            256,
	"mixed5a_3x3_bottleneck_pre_relu/conv": 160,
	"mixed5a_3x3_pre_relu/conv":            320,
	"mixed5a_5x5_bottleneck_pre_relu/conv": 48,
	"mixed5a_5x5_pre_relu/conv":            128,
	"mixed5a_pool_reduce_pre_relu/conv":    128,
	"mixed5b_1x1_pre_relu/conv":            384,
	"mixed5b_3x3_bottleneck_pre_relu/conv": 192,
	"mixed5b_3x3_pre_relu/conv":            384,
	"mixed5b_5x5_bottleneck_pre_relu/conv": 48,
	"mixed5b_5x5_pre_relu/conv":            128,
	"mixed5b_pool_reduce_pre_relu/conv":    128,
	"head0_bottleneck_pre_relu/conv":       128,
	"head1_bottleneck_pre_relu/conv":       128,
}
