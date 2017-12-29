package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/TableMountain/goydl"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
)

var isJob bool
var basePath string

// Truncate todo
func Truncate(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}

func now() string {
	return time.Now().Format(time.Kitchen)
}

// Dream is exported so it can be an api, haha what fun. Games perhaps? Stock trading? Some real time video effect?
func Dream(c *gin.Context) {
	start := time.Now()
	defer func() {
		elapsed := fmt.Sprintf("%s %s", now(), time.Since(start))
		elapsed = strings.Split(elapsed, ".")[0] + "s"
		Log.Info("job took ", elapsed)
		mel.Broadcast([]byte(elapsed))
	}()
	yt := c.PostForm("yt")
	fps := c.PostForm("fps")
	ov := c.PostForm("ov") //data the user uploaded we want
	ovf := c.PostForm("ovf")
	of := c.PostForm("of")
	oo := c.PostForm("oo")
	it := c.PostForm("iterations")
	oc := c.PostForm("octaves")
	la := c.PostForm("layer")
	rl := c.PostForm("rl")
	Log.Info("rl: ", rl)
	ow := c.PostForm("ow")
	li := c.PostForm("li")
	iw := c.PostForm("iw")
	rle := c.PostForm("rle")
	ocs := c.PostForm("ocscale")
	ch := c.PostForm("ch")
	// stretch:=c.Postform("stretchvideo")
	isJob = true
	defer func() {
		isJob = false
	}()

	Log.WithFields(logrus.Fields{
		"event": "new job started",
	})
	jobLog.WithFields(logrus.Fields{
		"time":  time.Now().UTC().UnixNano(),
		"title": name,
		"fps":   fps,
		"it":    it,
		"oc":    oc,
		"la":    la,
		"rl":    rl,
		"ow":    ow,
		"li":    li,
		"iw":    iw,
		"rle":   rle,
	})

	Log.Info("base path is ", basePath)

	newJobLog(name)
	//let's save interesting job metadata for the user in a tidy format (err logs, srv logs kept with the binary or maybe put in bind dir? wip)
	jobLog.WithFields(logrus.Fields{
		"fps":                         fps,
		"iterations":                  it,
		"octaves":                     oc,
		"layer":                       la,
		"linear increase":             li,
		"iteration waver":             iw,
		"octave waver":                ow,
		"randomization type":          rl,
		"random layer every n frames": rle,
	}).Info("job name: ", name)

	//
	var uploadedFile, framesDirPath string
	var name, fullName, ext string
	if yt != "" { //if "yt" checkbox checked
		youtubeDl := goydl.NewYoutubeDl()
		for { //we loop until we got an acceptable ytURL
			fmt.Println("waiting...")
			youtubeDl.VideoURL = ytURL
			fmt.Println("videoURL:", ytURL)
			if ytURL == "" { //we didn't get a url, so just cancel the job
				Log.Info("the url was blank (therefore no good ytURL yet), so just cancel the job")
				return
			}
			info, err := youtubeDl.GetInfo()
			if err != nil {
				Log.WithFields(logrus.Fields{
					"event": "ytdl",
					"error": err,
				}).Error("we should never fail here")
				continue
			}
			fmt.Println(youtubeDl.VideoURL, "blah")
			ext = info.Ext
			name = strings.Split(info.Title, " ")[0]
			fullName = name + ".mp4"
			if alreadyHave(basePath + "/frames/" + name) {
				name = renamer(name)
				fullName = name + ".mp4"
				Log.Info("\nwe renamed as: ", fullName)
			}
			uploadedFile := fmt.Sprintf("%s/frames/%s/%s.mp4", basePath, name, name)
			fmt.Println("uploaded file: ", uploadedFile)
			youtubeDl.Options.Output.Value = uploadedFile
			youtubeDl.Options.Format.Value = "mp4"
			cmd, err := youtubeDl.Download(youtubeDl.VideoURL)
			if err != nil {
				Log.WithFields(logrus.Fields{
					"event":        "error",
					"err":          err,
					"uploadedFile": uploadedFile,
				}).Error("dl'ing from yt failed w err")
			} else {
				Log.WithFields(logrus.Fields{
					"event": "download",
					"path":  uploadedFile,
				}).Info("downloaded a yt video")
				println("starting download")
				cmd.Wait()
				println("finished download")

				// make new folder for job
				framesDirPath = fmt.Sprintf("%s/frames/%s", basePath, name)
				if _, err := os.Stat(framesDirPath); os.IsNotExist(err) {
					if err = os.Mkdir(framesDirPath, 0777); err != nil {
						Log.Error("failed to make a new job dir w/ error: ", err)
					}
					Log.Info("frames folder for new job was created at ", framesDirPath)
				}

				break //we got our file, now we move on, we don't need to keep listening for URL
			}
		}
	} else { // if no youtube, then get file from form upload
		file, err := c.FormFile("file")
		if err != nil {
			Log.Error("failed to get file", err) //although this might not be an error as we support ytdl now
			return
		}
		name = strings.Split(file.Filename, ".")[0]
		fullName = file.Filename

		ext = strings.Split(fullName, ".")[1]
		if alreadyHave(basePath + "/frames/" + name) {
			name = renamer(name)
			fullName = name + "." + strings.Split(file.Filename, ".")[1]
			Log.Info("\nwe renamed as: ", fullName)
		}
		// make new folder for job
		framesDirPath = fmt.Sprintf("%s/frames/%s", basePath, name)
		if _, err := os.Stat(framesDirPath); os.IsNotExist(err) {
			if err = os.Mkdir(framesDirPath, 0777); err != nil {
				Log.Error("failed to make a new job dir w/ error: ", err)
			}
			Log.Info("frames folder for new job was created at ", framesDirPath)
		}
		uploadedFile = fmt.Sprintf("%s/%s", framesDirPath, fullName)
		if err := c.SaveUploadedFile(file, uploadedFile); err != nil {
			Log.Error("failed to save file at path ", uploadedFile, " err is: ", err)
		} else {
			Log.Info("saved file at path ", uploadedFile)
		}
	}
	// make a new output folder
	outputPath := fmt.Sprintf("%s/output", framesDirPath)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		os.Mkdir(outputPath, 0777)
		Log.Info("output folder for new job was created at ", outputPath)
	}
	Log.Info("saved output dir at path ", outputPath)
	uploadedFile = fmt.Sprintf("%s/%s", framesDirPath, fullName)
	mel.Broadcast([]byte(name))

	itsAVideo := false
	// decide what to do with the file we've gotten, if it's an image:
	if ext == "png" { //it's perfect, leave it alone...

	} else if ext == "jpg" || ext == "jpeg" {
		cmd, err := exec.Command("ffmpeg", "-i", uploadedFile, framesDirPath+"/"+name+".png").CombinedOutput()
		if err != nil {
			Log.Error("oops, failed trying to make some image of ext ", ext, " to png")
		} else {
			Log.Info("that's great, we got an image, those are easy, ffmpeg said:", string(cmd))
		}
	} else if ext == "gif" {
		itsAVideo = true
		Log.Info("trying to convert a gif")
		// ffmpeg -f gif -i giphy-downsized.gif  -pix_fmt yuv420p -c:v libx264 -movflags +faststart -filter:v crop='floor(in_w/2)*2:floor(in_h/2)*2' BAR.mp4
		savedMp4 := fmt.Sprintf("%s/frames/%s/%s.mp4", basePath, name, name)
		cmd := exec.Command("ffmpeg", "-f", "gif", "-i", uploadedFile, "-pix_fmt", "yuv420p", "-c:v", "libx264", "-movflags", "+faststart", "-filter:v", "crop='floor(in_w/2)*2:floor(in_h/2)*2'", savedMp4)
		cmd.Stdin = strings.NewReader("")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			Log.Error("failed to make mp4 from gif ", err)
		} else {
			uploadedFile = strings.Split(uploadedFile, ".")[0] + ".mp4"
			Log.Info("made mp4 from GIF")
		}
	} else { // if file not gif or img try to make it mp4
		itsAVideo = true
		Log.Info("ext: ", ext)
		Log.Info("file.filename ", fullName)
		if ext != "mp4" {
			cmd, err := exec.Command("ffmpeg", "-i", uploadedFile, strings.Split(uploadedFile, ".")[0]+".mp4").CombinedOutput()
			if err != nil {
				Log.Error("failed to make a .any to .mp4 , ", err)
			} else {
				Log.Info("made a ", ext, " into .mp4 with cmd ", string(cmd))
				err := os.Remove(uploadedFile)
				if err != nil {
					Log.Info("err removing original .mp4 as err: ", err)
				} else {
					uploadedFile = strings.Split(uploadedFile, ".")[0] + ".mp4"
					Log.Info("deleted original at ext: ", ext)
				}
			}
		}
	}
	// open finder
	if of == "of" {
		open.Run(framesDirPath)
	}
	if oo == "oo" {
		open.Run(outputPath)
	}

	if itsAVideo {
		// create  frames from mp4
		framesOut := fmt.Sprintf("%s/frames/%s/%s.png", basePath, name, "%d")
		Log.Info("framesOut: ", framesOut)
		cmd, err := exec.Command("ffmpeg", "-i", uploadedFile, "-vf", "fps="+fps, "-c:v", "png", framesOut).CombinedOutput()
		if err != nil {
			Log.Error("failed to make frames", err)
		} else {
			Log.Info("made frames from MP4 with cmd: ", string(cmd))
		}
	}
	Log.Info("ch is : ", ch)
	Log.Info("entering dreamer goroutine")
	// deep dream the frames
	cmd, err := exec.Command("python3", "folder.py", "--input", framesDirPath, "-ch", ch, "-os", ocs, "-it", it, "-oc", oc, "-la", la, "-rl", rl, "-rle", rle, "-li", li, "-iw", iw, "-ow", ow).CombinedOutput()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"event": "folder.py",
		}).Error("failed to dream", err)
		z := fmt.Sprintf("FAIL: python borked: %s", err.Error())
		mel.Broadcast([]byte(z))
	}
	Log.Info("done w/ dream loop, python said: ", string(cmd))
	if !itsAVideo {
		_, err := exec.Command("ffmpeg", "-i", outputPath+"/1.png", outputPath+"/1.jpg").CombinedOutput()
		if err != nil {
			Log.Error("failed to jpg the png", err)
		}
		return //if it's not a video, don't make an output.mp4
	}
	// put frames together into an mp4 in videos dir
	newVideo := fmt.Sprintf("%s/videos/%s", basePath, name+".mp4")
	frames := fmt.Sprintf("%s/output/%s.png", framesDirPath, "%d")
	Log.Info("frames to be turned into mp4 at: ", frames)
	// framesDir := fmt.Sprintf("%s/output/%s.png", framesDirPath, "%d")
	// ffmpeg -r 5 -f image2 -i '%d.png' -vcodec libx264 -crf 25 -pix_fmt yuv420p out.mp4
	cmd, err = exec.Command("ffmpeg", "-r", fps, "-f", "image2", "-i", frames, "-vcodec", "libx264", "-crf", "25", "-pix_fmt", "yuv420p", newVideo).CombinedOutput()
	if err != nil {
		Log.Error("still failing to output a video meh, ", err)
	} else {
		Log.Info("\nmade mp4 from frames")
	}
	if ov == "ov" {
		open.Run(basePath + "/videos")
	}

	//  is there sound?
	audio, err := exec.Command("ffprobe", uploadedFile, "-show_streams", "-select_streams", "a", "-loglevel", "error").CombinedOutput()
	if err != nil {
		Log.Error("Failed to test audio, ", err)
	}
	// add sound back in if there is any
	// ffmpeg -i 2171447000212516064.mp4 -i gold.mp4  -map 0:v -map 1:a output.mp4
	if len(audio) > 1 {
		Log.Info("there's sound in this clip")
		out, err := exec.Command("ffmpeg", "-y", "-i", newVideo, "-i", uploadedFile, "-map", "0:v", "-map", "1:a", basePath+"/videos/audio_"+name+".mp4").CombinedOutput()
		if err != nil {
			Log.Error("failed to add sound back", err)
		} else {
			Log.Info("fffmpeg added sound:", string(out))
			if ovf == "ovf" {
				open.Run(basePath + "/videos/audio_" + name + ".mp4")
			}
			// todo remove newVideo, so we only save one video, the one w/ audio
		}
	} else {
		Log.Info("there's no sound")
		if ovf == "ovf" {
			open.Run(newVideo)
		}
	}
	//stretch video enabled?
	// ffmpeg -i input.mp4 -vf scale=ih*16/9:ih,scale=iw:-2,setsar=1 -crf 20 -c:a copy YT.mp4
	// ffmpeg -i out.mp4 -vf scale=720x406,setdar=16:9 z.mp4
	// http://www.bugcodemaster.com/article/changing-resolution-video-using-ffmpeg
	// out, err := exec.Command("ffmpeg", "-y", "-i", newVideo, "-i", uploadedFile, "-map", "0:v", "-map", "1:a", basePath+"/videos/"+name+"_audio.mp4").CombinedOutput()
	// if err != nil {
	// 	Log.Error("failed to add sound back", err)
	// } else {
	// 	Log.Info("fffmpeg added sound:", string(out))
	// 	os.Remove(newVideo) //remove video w/o sound, we don't need it
	// 	if ovf == "ovf" {
	// 		open.Run(basePath + "/videos/" + name + "_audio.mp4")
	// 	}
	// 	// todo remove newVideo, so we only save one w/ audio
	// }

}
