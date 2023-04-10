package internals

import (
	"fmt"
	osUser "os/user"
)

func PrintUser() {
	user, err := osUser.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.Name)
	fmt.Println(user.Gid)
	fmt.Println(user.Uid)
	fmt.Println(user.Username)
}
