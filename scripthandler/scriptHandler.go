package sscripthandler

import (
	"fmt"
	"os/exec"
)

var shellPathOut string = fmt.Sprintf("%s%s", mainconfig.ScriptsLocation, "honeoutsidedoor.sh")
var shellPathIn string = fmt.Sprintf("%s%s", mainconfig.ScriptsLocation, "inside_door_hone.sh")

func OpenOutside() error {
	cmd := exec.Command(shellPathOut)
	return cmd.Run()
}

func OpenInside() error {
	cmd := exec.Command(shellPathIn)
	return cmd.Run()
}
