package main

import (
	"fmt"
	"os/user"
)

func main() {
	current, err := user.Current()
	if err != nil {
		fmt.Println("meh")
	}
	fmt.Println(current.Username)
	fmt.Println(current.Name)
	fmt.Println(current.HomeDir)
	fmt.Println(current.Uid)
	fmt.Println(current.Gid)
	ids, err := current.GroupIds()
	for item := range ids {
		fmt.Println(item)
	}
}
