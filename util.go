package main

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net"
	"os/exec"
	"os/user"
	"strings"
	"time"

	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/johntdyer/slackrus"
	log "github.com/sirupsen/logrus"
	filetype "gopkg.in/h2non/filetype.v1"
)

// Log is exported because log was taken already
var Log *log.Logger
var online chan (bool)

var currentUser string

func init() {
	go internet()
}

func internet() {
	for {
		timeout := time.Duration(1 * time.Minute)
		_, err := net.DialTimeout("tcp", "http://google.com", timeout)
		if err != nil {
			online <- false
			log.Info("Site unreachable, error: ", err)
		} else {
			online <- true
			log.Info("internet is up")
		}
	}
}

// makes sure we have our working dir to place all our files
func ensureDreamlyDirs() error {
	// usr, err := homedir.Dir()
	// if err != nil {
	// 	log.Error("failed to get homedir", err)
	// }
	// exp, err := homedir.Expand(usr)
	// if err != nil {
	// 	log.Error("failed to get expanded homedir", err)
	// }
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Info("dir of bind is ", dir)
	log.Info("Hello " + user.Name + " <3100LTC")
	log.Info("====")
	log.Info("Id: " + user.Uid)
	log.Info("Username: " + user.Username)
	log.Info("Home Dir: " + user.HomeDir)
	cmd, err := exec.Command("who").CombinedOutput()
	if err != nil {
		log.Error("failed to make frames", err)
	}
	currentUser = strings.Split(string(cmd), " ")[0]
	log.Info("this is the normal user: ", currentUser)
	basePath := fmt.Sprintf("/Users/%s/Desktop/dreamly", currentUser)
	// log.Info("dir: ", usr, " and expanded dir: ", exp, " and basePath to be working from is ", basePath)
	log.Info("dreamly folder  basePath: ", basePath)
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		err := os.Mkdir(basePath, 0777)
		if err != nil {
			log.Error("failed os.Mkdir", err)
		}
		log.Info("dreamly FOLDER was created")
	}
	// also make dirs that live inside it
	frames := fmt.Sprintf("%s/frames", basePath)
	if _, err := os.Stat(frames); os.IsNotExist(err) {
		err := os.Mkdir(frames, 0777)
		if err != nil {
			log.Error("failed os.Mkdir", err)
		}
		log.Info("dreamly frames was created")
	}
	audio := fmt.Sprintf("%s/audio", basePath)
	if _, err := os.Stat(audio); os.IsNotExist(err) {
		err := os.Mkdir(audio, 0777)
		if err != nil {
			log.Error("failed os.Mkdir", err)
		}
		log.Info("dreamly audio was created")
	}
	videos := fmt.Sprintf("%s/videos", basePath)
	if _, err := os.Stat(videos); os.IsNotExist(err) {
		err := os.Mkdir(videos, 0777)
		if err != nil {
			log.Error("failed os.Mkdir", err)
		}
		log.Info("dreamly video was created")
	}
	return nil
}

// saveFile will save whatever file it can
func saveFile(c *gin.Context) (string, string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Info("failed to get file", err)
	}
	name := strings.Split(file.Filename, ".")[0]
	path := fmt.Sprintf("$HOME/Desktop/%s", name)
	if alreadyHave(path) {
		name = renamer()
		path = fmt.Sprintf("$HOME/Desktop/%s", name)
		log.Info("\nwe renamed as: ", name)
	}
	// updateUser(name, c)

	// save the file
	savedFile := fmt.Sprintf("frames/%s/%s", name, file.Filename)
	if err := c.SaveUploadedFile(file, savedFile); err != nil {
		// c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		log.Fatal("failed to save", err)
	}
	return "", "", errors.New("failed to save file")
}

// checkFile sort out what kind of file it is and if we support it, else error
func checkFile(c *gin.Context) (string, error) {
	reader, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Info("can't get file from form", err)
		c.String(200, "file missing from upload, please try again with a file ")
		return "", errors.New("file missing from upload")
	}
	// check if it's really a gif
	fmt.Print("checking what kind of file")
	// log.Println(uploadFile.Filename)
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Print("not buffered")
	}
	kind, unknown := filetype.Match(buf)
	if unknown != nil {
		log.Info("Unknown file type: %s", unknown)
		return "", errors.New("bad file type, I can't work with it: ,")
	} else if kind.Extension == "video/mp4" {
		log.Info("it's a video!, todo: implement")
		return "video/mp4", nil
	} else if kind.Extension == "image/gif" {
		return "image/gif", nil
	} else if kind.Extension == "image/jpg" {
		return "image/jpg", nil
	} else {
		return "", errors.New("don't knwo what file this is ")
	}
}

// howManyOf returns list of .ext at a path
func howManyOf(ext string, pathS string) int {
	list := make([]string, 0, 100)

	// get all gifs in deepgif dir
	err := filepath.Walk(pathS, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ext { //like .mp4 or .mov or .png 
			log.Info(path)
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		log.Infof("walk error [%v]\n", err)
	}
	return len(list)
}

// deepGIFFIles returns list of gifs in dir deepgif
func deepGIFFiles() []string {
	list := make([]string, 0, 100)

	// get all gifs in deepgif dir
	err := filepath.Walk("deepgif", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".mp4" {
			log.Info(path)
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return list
}

func alreadyHave(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
func renamer() string {
	fmt.Print("\n We had to rename the file")
	return fmt.Sprintf("%d", rand.Int())
}

func newLogger() *log.Logger {

	// if Log != nil {
	// 	return Log
	// }
	// writer, err := rotatelogs.New(
	// 	"logrus.%Y%m%d%H%M",
	// 	rotatelogs.WithLinkName("logrus"),
	// )
	// if err != nil {
	// 	Log.Error("failed to setup logs", err)
	// }
	// log.AddHook(lfshook.NewHook(lfshook.WriterMap{
	// 	log.InfoLevel:  writer,
	// 	log.ErrorLevel: writer,
	// }, &log.JSONFormatter{}))

	// only send logs to slack if we're online
	go func() {
		for {
			select {
			case online <- true:
				log.Info("adding slackrus hook")
				log.AddHook(&slackrus.SlackrusHook{
					HookURL:        "https://hooks.slack.com/services/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",
					AcceptedLevels: slackrus.LevelThreshold(log.DebugLevel),
					Channel:        "#test",
					IconEmoji:      ":ghost:",
					Username:       "testbot",
				})
			}
		}
	}()

	return Log
}
