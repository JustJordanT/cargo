package main

import (
	"fmt"
	osUser "os/user"
)

func main() {
	printUser()
}

func printUser() {
	user := osUser.Current()
	fmt.Println(user)
}
