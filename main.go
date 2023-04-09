package main

import (
	"fmt"

	"github.com/justjordant/cargo/internals"
)

func main() {
	fmt.Println(internals.GetCrateUrl("cli-ops.yml"))
	// TODO Parse need to be cleaned up
	// TODO Need to pass in what file we want to download from the `Crates` to then be installed down the pipe
}
