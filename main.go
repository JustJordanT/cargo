package main

import (
	"fmt"

	"github.com/justjordant/cargo/cmd"
	"github.com/justjordant/cargo/internals"
	"github.com/justjordant/cargo/internals/initUtils"
)

func main() {

	// TODO - This section might be better off in BASH, where we can pull in the all the needed this.
	createFile := false
	if createFile {
		initUtils.InitDirs()
	}
	cmd.Execute()

	url, fileName := internals.GetCrateUrl("tool.yml")
	fmt.Println(url)
	fmt.Println(internals.DownloadCargoYaml(url, internals.FormatFileName(fileName)))
	// TODO Parse need to be cleaned up
	// TODO Need to pass in what file we want to download from the `Crates` to then be installed down the pipe
}
