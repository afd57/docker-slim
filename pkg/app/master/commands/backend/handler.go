package backend

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/docker-slim/docker-slim/pkg/app"
	//"github.com/docker-slim/docker-slim/pkg/app/master/commands"
)

type ovars = app.OutVars

func runBackend() {
	destinationDirPath, err := os.Getwd()
	//script := filepath.Join(destinationDirPath, "pf_file_monitor.sh")
	log := filepath.Join(destinationDirPath, "nohup2.out")
	cmd := exec.Command("nohup", "sh", "-c", "/Users/afdagli/Desktop/Study/docker-slim/bin/mac/docker-slim-backend")

	f, err := os.Create(log)
	if err != nil {
		// handle error
	}

	// redirect both stdout and stderr to the log file
	cmd.Stdout = f
	cmd.Stderr = f

	// start command (and let it run in the background)
	err = cmd.Start()
	if err != nil {
		// handle error
	}

}

func shutdownBackend() {
	cmd := exec.Command("pkill", "docker-slim-backend")

	// start command (and let it run in the background)
	err := cmd.Start()
	fmt.Println(err)
}

// OnCommand implements the 'backend' docker-slim command
func OnCommand(backendEnable bool) {
	fmt.Println(backendEnable)
	if backendEnable {
		fmt.Println("Backend Enable")
		runBackend()
	} else {
		fmt.Println("Backend Disable")
		shutdownBackend()
	}

}
