package main

import (
	"fmt"

	"github.com/lulascoca/gatetool/controlServer"
)

func main() {
	fmt.Printf("yeah, just started listening for new connections\n")

	controlServer.ListenMain("0.0.0.0:8000")
}
