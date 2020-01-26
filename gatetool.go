package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

const shell_path_out string = "./scripts_waves/honeoutsidedoor.sh"
const shell_path_in string = "./scripts_waves/inside_door_hone.sh"

func open_inside() error {
	cmd := exec.Command(shell_path_out)
	err := cmd.Run()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func open_outside() error {
	cmd := exec.Command(shell_path_in)
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
	access_key, _ := ioutil.ReadFile("./passwd")
	actual_pass, err := bcrypt.GenerateFromPassword([]byte(strings.Trim(string(access_key), " \n")), 10)

	// password from user
	passwd := r.FormValue("pass")

	if err = bcrypt.CompareHashAndPassword(actual_pass, []byte(passwd)); err != nil {
		w.WriteHeader(69)
		fmt.Printf("nope\n")
		return
	} else {
		err := open_inside()
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

	// the actual password thats read from a file called passwd in the same directory
	access_key, _ := ioutil.ReadFile("./passwd")
	actual_pass, err := bcrypt.GenerateFromPassword([]byte(strings.Trim(string(access_key), " \n")), 10)

	// password from user
	passwd := r.FormValue("pass")

	if err = bcrypt.CompareHashAndPassword(actual_pass, []byte(passwd)); err != nil {
		w.WriteHeader(69)
		fmt.Fprintf(w, "nope\n")
		return
	} else {
		err := open_outside()
		if err != nil {
			log.Fatal(err)
			fmt.Fprintf(w, err.Error())
			return
		}
		fmt.Fprintf(w, "called outside successfully.\n")
	}
}

func main() {
	// creates new router
	router := mux.NewRouter().StrictSlash(false)

	// this section handles all the parts of the thing
	router.HandleFunc("/", baseAccess)
	router.HandleFunc("/api/out", apiAccessOut).Methods("POST")
	router.HandleFunc("/api/in", apiAccessIn).Methods("POST")
	router.HandleFunc("/api/out/", apiAccessOut).Methods("POST")
	router.HandleFunc("/api/in/", apiAccessIn).Methods("POST")

	log.Fatal(http.ListenAndServe("0.0.0.0:8001", router))
}
