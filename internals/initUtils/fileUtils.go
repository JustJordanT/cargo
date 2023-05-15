package initUtils

import (
	"errors"
	"log"
	"os"
)

var CargoPath = "/usr/local/cargo/"                    // Specify the directory path here
var ContainerPath = "/usr/local/cargo/container"       // Specify the directory path here
var CratePath = "/usr/local/cargo/crates"              // Specify the directory path here
var LogsPath = "/usr/local/var/log/cargo/loggious.txt" // Specify the directory path here

func CheckDirExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrExist)
}

func InitDirs() {
	// Use os.MkdirAll to create the directory along with any necessary parent directories
	createDir(CargoPath)
	createDir(LogsPath)
	createDir(ContainerPath)
}

func createDir(dirPath string) {
	//  fmt.Println("CheckDirNotExist..." + strconv.FormatBool(CheckDirExist(dirPath))) TODO make this into a log
	//  fmt.Println("Creating directory...." + dirPath) TODO make this into a log
	if CheckDirExist(dirPath) {
		err := os.Mkdir(dirPath, 0755) // 0755 is the default permission for a new directory
		if err != nil {
			// Handle error if any
			log.Fatal("Failed to create directory:", err)
			return
		}
	}
}
