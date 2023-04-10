package main

import (
	"github.com/JustJordanT/cargo/cmd"
	"github.com/justjordant/cargo/internals/initUtils"
)

func main() {

	initUtils.InitDirs()
	cmd.Execute()

	//  url := internals.GetCrateUrl("example.yml")
	//  fmt.Println(url)
	//  fmt.Println(internals.DownloadCargoYaml(fileName, url))
	// TODO Parse need to be cleaned up
	// TODO Need to pass in what file we want to download from the `Crates` to then be installed down the pipe
}
