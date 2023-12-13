package main

import (
    "fmt"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
	"encoding/json"
)

// Handle GET requests at /scripts endpoint
func getScriptsHandler(w http.ResponseWriter, r *http.Request) {
    // use mainConfig.ScriptsLocation to get the scripts location
	var scriptNames []string
    for _, script := range MainConfig.Scripts {
        scriptNames = append(scriptNames, script.Name)
    }
    scriptNamesJson, err := json.Marshal(scriptNames)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(scriptNamesJson)
}

func runScriptHandler(w http.ResponseWriter, r *http.Request) {
    scriptName := r.FormValue("script")

    // Find the script in the MainConfig
    var scriptFileName string
    for _, script := range MainConfig.Scripts {
        if script.Name == scriptName {
            scriptFileName = script.FileName
            break
        }
    }

    if scriptFileName == "" {
        http.Error(w, "Script not found", http.StatusNotFound)
        return
    }

    // Check if script exists
    scriptPath := filepath.Join(MainConfig.ScriptsLocation, scriptFileName)
    if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
        http.Error(w, "Script file not found", http.StatusNotFound)
        return
    }

    // Run the script
    cmd := exec.Command("/bin/sh", scriptPath)
    if err := cmd.Run(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Script %s ran successfully.\n", scriptName)
}
