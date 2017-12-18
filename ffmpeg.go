package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"os/exec"
// 	"path/filepath"

// 	log "github.com/sirupsen/logrus"
// )

// var CWD string

// func init() {
// 	CWD, err := filepath.Abs(filepath.Dir(os.Args[0]))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(CWD)
// }

// func ffmpeg() (string, error) {
// 	cmd, err := exec.LookPath("ffmpeg")
// 	if err != nil {
// 		log.Info("can't find ffmpeg")
// 	} else {
// 		return "ffmpeg", nil
// 	}
// 	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
// 		// path/to/whatever does not exist
// 	}
// 	out, err := os.Create("output.txt")
// 	defer out.Close()

// 	resp, err := http.Get("http://example.com/")
// 	defer resp.Body.Close()

// 	n, err := io.Copy(out, resp.Body)
// 	return ("./ffmpeg", nil)
// }
