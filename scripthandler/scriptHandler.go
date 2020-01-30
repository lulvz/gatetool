package sscripthandler

import (
	"fmt"
	"os/exec"
)

// var shellPathOut string = fmt.Sprintf("%s%s", mainconfig.ScriptsLocation, "honeoutsidedoor.sh")
// var shellPathIn string = fmt.Sprintf("%s%s", mainconfig.ScriptsLocation, "inside_door_hone.sh")

func OpenOutside(ScriptsLocation string) error {
	var shellPathOut string = fmt.Sprintf("%s%s", mainconfig.ScriptsLocation, "honeoutsidedoor.sh")

	cmd := exec.Command(shellPathOut)
	return cmd.Run()
}

func OpenInside(ScriptsLocation string) error {
	var shellPathIn string = fmt.Sprintf("%s%s", mainconfig.ScriptsLocation, "inside_door_hone.sh")

	cmd := exec.Command(shellPathIn)
	return cmd.Run()
}
