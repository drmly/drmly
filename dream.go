package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/skratchdot/open-golang/open"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	// https://github.com/syncthing/syncthing
)

var basePath string

func mp4ToDream(c *gin.Context) {
	log.Info("mp42dream")
	// usr, err := homedir.Dir()
	// if err != nil {
	// 	log.Error("failed to get homedir", err)
	// }
	// exp, err := homedir.Expand(usr)
	// if err != nil {
	// 	log.Error("failed to get expanded homedir", err)
	// }
	// get the file
	file, err := c.FormFile("file")
	if err != nil {
		log.Info("failed to get file", err)
		c.String(200, "no file added huh? ", err)
	}
	name := strings.Split(file.Filename, ".")[0]
	fullName := file.Filename
	// basePath is dreamly stores stuff here on desktop
	basePath = fmt.Sprintf("/Users/%s/Desktop/dreamly", currentUser)
	log.Info("base path is ", basePath)
	if alreadyHave(basePath) {
		name = renamer()
		fullName = name + "." + strings.Split(file.Filename, ".")[1]
		log.Info("\nwe renamed as: ", fullName)
	}
	// make new folder for job
	framesDirPath := fmt.Sprintf("%s/frames/%s", basePath, name)
	if _, err := os.Stat(framesDirPath); os.IsNotExist(err) {
		if err = os.Mkdir(framesDirPath, 0777); err != nil {
			log.Error("failed to make a new job dir w/ error: ", err)
		}
		log.Info("frames folder for new job was created at ", framesDirPath)
	}

	// make a new output folder
	outputPath := fmt.Sprintf("%s/output", framesDirPath)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		os.Mkdir(outputPath, 0777)
		log.Info("output folder for new job was created at ", outputPath)
	}
	log.Info("saved output dir at path ", outputPath)

	// save the file
	savedFilePath := fmt.Sprintf("%s/%s", framesDirPath, fullName)
	if err := c.SaveUploadedFile(file, savedFilePath); err != nil {
		log.Error("failed to save file at path ", savedFilePath, " err is: ", err)
	}
	log.Info("saved file at path ", savedFilePath)
	// if file not mp4 try to make it mp4
	ext := strings.Split(file.Filename, ".")[1]
	log.Info("ext: ", ext)
	log.Info("file.filename ", file.Filename)
	if ext != "mp4" {
		cmd,err := exec.Command("ffmpeg", "-i", savedFilePath, strings.Split(savedFilePath,".")[0]+".mp4").CombinedOutput()
		if err != nil {
			log.Error("failed to make an .any to .mp4 , " ,err)
			} else {
				log.Info("made a ",ext ," into .mp4 with cmd " , string(cmd))
				err := os.Remove(savedFilePath)
				if err != nil {
					log.Info("err removing original .mp4 as err: ", err)
					c.String(200,"failed with file extension ", ext, " \n try again with better supported file, like mp4")
					return
				}
				savedFilePath = strings.Split(savedFilePath,".")[0]+".mp4"
				log.Info("deleted original at ext: ", ext)
			}
		}
	// open finder
	open.Run(framesDirPath)
	open.Run(outputPath)

	// create  frames from mp4
	framesOut := fmt.Sprintf("%s/frames/%s/%s.png", basePath, name, "%d")
	log.Info("framesOut: ", framesOut)
	// fps := c.PostForm("fps")
	cmd, err := exec.Command("ffmpeg", "-i", savedFilePath, "-vf", "fps=5", "-c:v", "png", framesOut).CombinedOutput()
	if err != nil {
		log.Error("failed to make frames", err)
	} else {
		log.Info("made frames from MP4 with cmd: ", string(cmd))
	}
	// deep dream the frames
	log.Info("inside dreamer loop")
	it := c.PostForm("iterations")
	oc := c.PostForm("octaves")
	la := c.PostForm("layer")
	rl := c.PostForm("rl")
	ow:= c.PostForm("ow")
 	li := c.PostForm("li")
	iw := c.PostForm("iw")
	log.Info("fruckkkk")
	cmd, err = exec.Command("python3","folder.py", "--input", framesDirPath,  "-it", it, "-oc", oc, "-la", la, "-rl", rl, "-li", li, "-iw", iw, "-ow", ow  ).CombinedOutput()
	if err != nil {
		log.Error("failed to dream", err)
		c.String(200, "Abort, this app is crashing, can't dream, probs the platform ur using is not OSX Sierra or youre not starting the app from terminal")
	}
	log.Info("done w/ dream loop, python said: ", string(cmd))

	// put frames together into an mp4 in videos dir
	newVideo := fmt.Sprintf("%s/videos/%s", basePath, name+".mp4")
	pathOk := func(p string) error {
		if _, err := os.Stat(p); err != nil {
			return nil
		} else {
			p = fmt.Sprintf("%s/%s.%s", basePath, renamer(), strings.Split(file.Filename, ".")[1])
			log.Info("new video to be made already in vidoe dir, renamed to :", p)
			return err
		}
	}
	for {
		err = pathOk(newVideo)
		if err != nil {
			pathOk(newVideo)
		} else {
			log.Info("new video to be made at ", newVideo)
			break
		}
	}

	frames := fmt.Sprintf("%s/output/%s.png", framesDirPath, "%d")
	log.Info("frames to be turned into mp4 at: ", frames)
	// framesDir := fmt.Sprintf("%s/output/%s.png", framesDirPath, "%d")
	_, err = exec.Command("ffmpeg", "-r", "5", "-f", "image2", "-i", frames, "-vcodec", "libx264", "-crf", "25", "-pix_fmt", "yuv420p", newVideo).CombinedOutput()
	if err != nil {
		log.Error("still failing to output a video meh, ", err)
	} else {
		log.Info("\nmade mp4 from frames")
	}

	open.Run(basePath + "/videos")
	open.Run(newVideo)

	//  is there sound?
	audio, err := exec.Command("ffprobe", savedFilePath, "-show_streams", "-select_streams", "a", "-loglevel", "error").CombinedOutput()
	if err != nil {
		log.Error("Failed to test audio, ", err)
	}
	// add sound back in if there is any
	// ffmpeg -i 2171447000212516064.mp4 -i gold.mp4  -map 0:v -map 1:a output.mp4 
	if len(audio) > 1 {
		log.Info("there's sound in this clip")
		out, err := exec.Command("ffmpeg", "-y", "-i", newVideo, "-i", savedFilePath, "-map" ,"0:v" , "-map", "1:a", basePath+"/videos/"+name+".mp4").CombinedOutput()
		if err != nil {
			log.Error("failed to add sound back", err)
		}
		log.Info("fffmpeg tried to add sound:", string(out))
	} else {

		log.Info("there's no sound")
	}
}
