package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

// gets config file name from env variable
var webConfigFile string = os.Getenv("GATEWEBCFG")
var mainConfigFile string = os.Getenv("GATECFG")
var WebConfig webConfig
var MainConfig mainConfig

func init() {
    var debugaa bool = true
    if debugaa {
        webConfigFile = "./webConfigDebug.json"
        mainConfigFile = "./configDebug.json"
    }
    WebConfig = readWebConfig(webConfigFile)
    MainConfig = readConfig(mainConfigFile)
}


func baseAccess(w http.ResponseWriter, r *http.Request) {
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

func main() {
    // base server handling
    http.HandleFunc("/", baseAccess)
    // api handling
    http.HandleFunc("/api/scripts", getScriptsHandler)
    http.HandleFunc("/api/run-script", runScriptHandler)
    // concatenate the address and the port into one string to pass to the handler
    fulladdr := fmt.Sprintf("%s:%d", WebConfig.Address, WebConfig.Port)
    // log to know the server has started and return error if necessary
    fmt.Printf("starting to listen on port %d\n", WebConfig.Port)
    log.Fatal(http.ListenAndServe(fulladdr, nil))
}
