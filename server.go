package main

import (
	"fmt"
	"os/user"

	"github.com/sckelemen/grpc/AppShell"
	"github.com/sckelemen/grpc/UpdateManager"
)

func main() {
	UpdateManager, err := UpdateManager.New("https://localhost/UpdateManager")
	if err != nil {
		fmt.Println(err)
		return
	}
	UpdateManager.CheckForUpdates()
	AppShell := AppShell.ApplicationShell{}
	AppShell.LoadApplication("test")

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
