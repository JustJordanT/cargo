package main

import (
	"github.com/justjordant/cargo/cmd"
	"github.com/justjordant/cargo/internals/initUtils"
	"github.com/justjordant/cargo/internals/loggy"
)

func main() {

	loggious := loggy.New()

	// TODO - This section might be better off in BASH, where we can pull in the all the needed this.
	// TODO - Don't worry about the code sigining thing for the application,
	// focus on installing of 3rd party apps first and then circle back on this down the road.
	// TODO - This should not be here, this should be moved into the installer repo. -->
	createFile := false
	if createFile {
		initUtils.InitDirs()
	}
	cmd.Execute()
	loggious.Log(0, "Cargo Info...")
	loggious.Log(1, "Cargo Fatal error...")
	loggious.Log(2, "Cargo Error...")
	loggious.Log(3, "Cargo Warn...")

	// TODO Parse need to be cleaned up
	// TODO Need to pass in what file we want to download from the `Crates` to then be installed down the pipe
}
