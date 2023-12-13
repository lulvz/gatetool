package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

bool debug = true

// gets config file name from env variable
var webConfigFile string = "./webConfigDebug.json"
var mainConfigFile string = "./configDebug.json"

if debug {
	webConfigFile = os.Getenv("GATEWEBCFG")
	mainConfigFile = os.Getenv("GATECFG")
}
var webconfig webConfig = readWebConfig(webConfigFile)
var mainconfig mainConfig = readConfig(mainConfigFile)

func baseAccess(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("heh\n")
	fmt.Fprintf(w, "pls")
}

func main() {
	// base server handling
	http.HandleFunc("/", baseAccess)

	// api handling
	http.HandleFunc("/api/scripts", getScriptsHandler)
	http.HandleFunc("/api/run-script", runScriptHandler)

	// concatenate the address and the port into one string to pass to the handler
	fulladdr := fmt.Sprintf("%s:%d", webconfig.Address, webconfig.Port)

	// log to know the server has started and return error if necessary
	fmt.Printf("starting to listen on port %d\n", webconfig.Port)
	log.Fatal(http.ListenAndServe(fulladdr, nil))
}

func main() {
    http.HandleFunc("/scripts", getScriptsHandler)
    http.HandleFunc("/run-script", runScriptHandler)
    http.ListenAndServe(":8080", nil)
}