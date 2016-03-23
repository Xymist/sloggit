package main

import (
	"os"
	"os/exec"
	"runtime"
)

var clear = map[string]func(){
	"linux": func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
	"windows": func() {
		cmd := exec.Command("cls") //Windows example it is untested, but I think its working
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
}

func callClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}
