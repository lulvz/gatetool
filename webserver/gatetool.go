package webserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

const shellPathOut string = "/home/pi/gatetool/scripts_waves/honeoutsidedoor.sh"
const shellPathIn string = "/home/pi/gatetool/scripts_waves/inside_door_hone.sh"

func openInside() error {
	cmd := exec.Command(shellPathIn)
	err := cmd.Run()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func openOutside() error {
	cmd := exec.Command(shellPathOut)
	err := cmd.Run()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func baseAccess(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome lmao !")
}

func apiAccessIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "revamp bitch\n")

	// the actual password thats read from a file called passwd in the same directory
	accessKey, _ := ioutil.ReadFile("/home/pi/wavv/passwd")
	// password from user
	passwd := r.FormValue("pass")

	if strings.Trim(string(accessKey), "\n ") != passwd {
		fmt.Fprintf(w, "nope\n")
		return
	} else {
		err := openInside()
		if err != nil {
			log.Fatal(err)
			fmt.Fprintf(w, err.Error())
			return
		}
		fmt.Fprintf(w, "called inside successfully.\n")
	}
}

func apiAccessOut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "revamp bitch\n")

	// the actual password that reads from a file called passwd in the same directory
	accessKey, _ := ioutil.ReadFile("/home/pi/wavv/passwd")
	// password from user
	passwd := r.FormValue("pass")

	if strings.Trim(string(accessKey), "\n ") != passwd {
		fmt.Fprintf(w, "nope\n")
		return
	} else {
		err := openOutside()
		if err != nil {
			log.Fatal(err)
			fmt.Fprintf(w, err.Error())
			return
		}
		fmt.Fprintf(w, "called outside successfully.\n")
	}
}

func ListenMain(address string) {
	// creates new router
	router := mux.NewRouter().StrictSlash(false)
	fmt.Printf("started listening on port 8000\n")

	// this section handles all the parts of the thing
	router.HandleFunc("/", baseAccess)
	router.HandleFunc("/api/out", apiAccessOut).Methods("POST")
	router.HandleFunc("/api/in", apiAccessIn).Methods("POST")
	router.HandleFunc("/api/out/", apiAccessOut).Methods("POST")
	router.HandleFunc("/api/in/", apiAccessIn).Methods("POST")

	log.Fatal(http.ListenAndServe(address, router))
}
