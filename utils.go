package sysproxy

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const DEBUG = false

func SaveToFile(name string, data []byte, permission os.FileMode) {
	if DEBUG {
		fmt.Println("Saving to file", name)
	}
	if fileInfo, err := os.Stat(name); err == nil {		// file already exists
		if DEBUG {
			fmt.Println("File exists!")
		}
		size, mode := fileInfo.Size(), fileInfo.Mode()
		if size == int64(len(data)) && mode == permission {
			if DEBUG {
				fmt.Println("File unchanged, skip.")
			}
			return
		}
	}
	// write to file
	err := os.WriteFile(name, data, permission)
	if err != nil {
		log.Fatal(err)
	}
}

func RunCmd(cmd string, args ...string) string {
	if DEBUG {
		fmt.Println("Running cmd", cmd, args)
	}
	outputBytes, err := exec.Command(cmd, args...).Output()
	output := strings.TrimSpace(string(outputBytes))
	if err != nil {
		log.Fatalf("$ %s\n%s\n%s", cmd, output, err)
	}
	return output
}

func RunBashCmd(cmd string) string {
	if DEBUG {
		fmt.Println("Running bash cmd", cmd)
	}
	outputBytes, err := exec.Command("/bin/bash", "-c", cmd).Output()
	output := strings.TrimSpace(string(outputBytes))
	if err != nil {
		log.Fatalf("$ %s\n%s\n%s", cmd, output, err)
	}
	return output
}
