package main

import (
	"fmt"

	"github.com/lulascoca/gatetool/controlserver"
)

func main() {
	fmt.Printf("yeah, just started listening for new connections\n")

	controlserver.ListenMain("0.0.0.0:8000")
}
