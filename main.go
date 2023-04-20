package main

import (
	"github.com/justjordant/cargo/cmd"
	"github.com/justjordant/cargo/internals/initUtils"
)

func main() {

	// TODO - This section might be better off in BASH, where we can pull in the all the needed this.
	// TODO - Don't worry about the code sigining thing for the application,
	// focus on installing of 3rd party apps first and then circle back on this down the road.
	createFile := false
	if createFile {
		initUtils.InitDirs()
	}
	cmd.Execute()

	// TODO Parse need to be cleaned up
	// TODO Need to pass in what file we want to download from the `Crates` to then be installed down the pipe
}
