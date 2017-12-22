package main

import (
	"fmt"
	"os"

	"github.com/takama/daemon"
)

func main() {
	srv, err := daemon.New(name, description, dependencies...)
	if err != nil {
		Log.Error("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		Log.Error(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
