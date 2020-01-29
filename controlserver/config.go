package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var homeDir string = os.Getenv("HOME")

// web Config Info from config file
type webConfig struct {
	// webserver constants
	PasswdLocation string

	Address string
	Port    int
}

// main Config Info from config file
type mainConfig struct {
	// main app constants
	ScriptsLocation string
}

// WEB CONFIG FUNCS

// reads the hardcoded default config file
func readWebDefaultConfig() webConfig {
	// create new config object of type Config
	var config webConfig

	// open default file and log if error
	var defaultPath = fmt.Sprintf("%s/%s", homeDir, ".gateToolFiles/gateToolWebConfig.json")
	jsonFileDat, err := os.Open(defaultPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("### opened default file ###\n\n")
	// convert opened file to bytearray
	byteValue, _ := ioutil.ReadAll(jsonFileDat)

	// open as json finally
	json.Unmarshal(byteValue, &config)

	defer jsonFileDat.Close()
	return config
}

// ReadConfig Reads info from config file
func readWebConfig(file string) webConfig {
	if len(file) != 0 {
		// create new config object of type Config
		var config webConfig

		// try to open file and log if error
		jsonFileDat, err := os.Open(file)
		fmt.Printf("%s\n", file)
		if err != nil {
			log.Fatal(err)
		}
		// inform file was opened from env variable instead of default location
		fmt.Printf("### opened file from env ###\n\n")

		// convert opened file to bytearray
		byteValue, _ := ioutil.ReadAll(jsonFileDat)

		// open as json finally
		json.Unmarshal(byteValue, &config)

		defer jsonFileDat.Close()
		return config
	}
	// in case the env var is no defined just use the default config
	return readWebDefaultConfig()
}

// MAIN CONFIG FUNCS

// reads the hardcoded default config file
func readDefaultConfig() mainConfig {
	// create new config object of type Config
	var config mainConfig

	// open default file and log if error
	var defaultPath = fmt.Sprintf("%s/%s", homeDir, ".gateToolFiles/gateToolConfig.json")
	jsonFileDat, err := os.Open(defaultPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("### opened default file ###\n\n")
	// convert opened file to bytearray
	byteValue, _ := ioutil.ReadAll(jsonFileDat)

	// open as json finally
	json.Unmarshal(byteValue, &config)

	defer jsonFileDat.Close()
	return config
}

// ReadConfig Reads info from config file
func readConfig(file string) mainConfig {
	if len(file) != 0 {
		// create new config object of type Config
		var config mainConfig

		// try to open file and log if error
		jsonFileDat, err := os.Open(file)
		fmt.Printf("%s\n%d\n", file, len(file))

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("### opened file from env ###\n\n")
		// convert opened file to bytearray
		byteValue, _ := ioutil.ReadAll(jsonFileDat)

		// open as json finally
		json.Unmarshal(byteValue, &config)

		defer jsonFileDat.Close()
		return config
	}
	// in case the env var is no defined just use the default config
	return readDefaultConfig()
}
