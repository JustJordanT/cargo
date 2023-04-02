package main

import "github.com/justjordant/cargo/internals"

func main() {
	internals.ParseCargo()
	// TODO Parse need to be cleaned up
	// TODO Need to pass in what file we want to download from the `Crates` to then be installed down the pipe
}
