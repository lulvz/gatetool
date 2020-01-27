package main

import (
	"fmt"

	"github.com/lulascoca/gatetool_go/webserver"
)

func main() {
	fmt.Printf("yeah, just started listening for new connections\n")

	webserver.ListenMain()
}
