package utilities

import (
	"os"
	"os/exec"
	"runtime"
)

//create a map for storing clear functions
var clear map[string]func()

func init() {
	//Initialize it
	clear = make(map[string]func())
	//Linux clear commands
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	//Windows clear command
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	//MAC clear command
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

func CallClear() {
	//runtime.GOOS -> linux, windows, darwin etc.
	value, ok := clear[runtime.GOOS]
	if ok { //if the platform is known
		value() //we execute it
	} else { //unsupported platform
		// Don't do anything
	}
}
