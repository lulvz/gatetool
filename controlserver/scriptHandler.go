package main

import (
	"fmt"
	"os/exec"
)

var shellPathOut string = fmt.Sprintf("%s%s", mainconfig.ScriptsLocation, "honeoutsidedoor.sh")
var shellPathIn string = fmt.Sprintf("%s%s", mainconfig.ScriptsLocation, "inside_door_hone.sh")

func openOutside() error {
	cmd := exec.Command(shellPathOut)
	return cmd.Run()
}

func openInside() error {
	cmd := exec.Command(shellPathIn)
	return cmd.Run()
}
