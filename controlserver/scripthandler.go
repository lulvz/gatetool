package main

import (
	"fmt"
	"os/exec"
	"io/ioutil"
)

// var shellPathOut string = fmt.Sprintf("%s%s", mainconfig.ScriptsLocation, "honeoutsidedoor.sh")
// var shellPathIn string = fmt.Sprintf("%s%s", mainconfig.ScriptsLocation, "inside_door_hone.sh")

func listScripts(ScriptsLocation string) ([]string, error) {
	files, err := ioutil.ReadDir(ScriptsLocation)
	if err != nil {
		return nil, err
	}

	var scripts []string
	for _, file := range files {
		scripts = append(scripts, file.Name())
	}

	return scripts, nil
}

func openScript(ScriptsLocation string, scriptName string) error {
	shellPath := fmt.Sprintf("%s/%s", ScriptsLocation, scriptName)

	cmd := exec.Command(shellPath)
	return cmd.Run()
}
