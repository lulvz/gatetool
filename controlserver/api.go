package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
)

// Handle GET requests at /scripts endpoint
func getScriptsHandler(w http.ResponseWriter, r *http.Request) {
    files, err := ioutil.ReadDir("./scripts")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var scripts []string
    for _, file := range files {
        scripts = append(scripts, file.Name())
    }

    fmt.Fprint(w, scripts)
}

// Handle POST requests at /run-script endpoint
func runScriptHandler(w http.ResponseWriter, r *http.Request) {
    scriptName := r.FormValue("script")

    // Check if script exists
    scriptPath := filepath.Join("./scripts", scriptName)
    if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
        http.Error(w, "Script not found", http.StatusNotFound)
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
