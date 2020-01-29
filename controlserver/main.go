package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// gets config file name from env variable
var webConfigFile string = os.Getenv("GATEWEBCFG")
var mainConfigFile string = os.Getenv("GATECFG")
var webconfig webConfig = readwebConfig(webConfigFile)
var mainconfig mainConfig = readConfig(mainConfigFile)

func baseAccess(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("heh\n")
	fmt.Fprintf(w, "pls")
}

func main() {
	// base server handling
	http.HandleFunc("/", baseAccess)

	// api handling
	http.HandleFunc("/api", apiAccess)
	http.HandleFunc("/api/", apiAccess)

	// concatenate the address and the port into one string to pass to the handler
	fulladdr := fmt.Sprintf("%s:%d", webconfig.Address, webconfig.Port)

	// log to know the server has started and return error if necessary
	fmt.Printf("starting to listen on port %d\n", webconfig.Port)
	log.Fatal(http.ListenAndServe(fulladdr, nil))
}
