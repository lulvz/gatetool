package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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
func readwebDefaultConfig() webConfig {
	// create new config object of type Config
	var config webConfig

	// open default file and log if error
	jsonFileDat, err := os.Open("/usr/local/gateTool/gateToolWeb/gateToolConfig.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("opened default file\n")
	// convert opened file to bytearray
	byteValue, _ := ioutil.ReadAll(jsonFileDat)

	// open as json finally
	json.Unmarshal(byteValue, &config)

	defer jsonFileDat.Close()
	return config
}

// ReadConfig Reads info from config file
func readwebConfig(file string) webConfig {
	if len(file) != 0 {
		// create new config object of type Config
		var config webConfig

		// try to open file and log if error
		jsonFileDat, err := os.Open(file)
		fmt.Printf("%s\n%d\n", file, len(file))

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("opened file from env\n")
		// convert opened file to bytearray
		byteValue, _ := ioutil.ReadAll(jsonFileDat)

		// open as json finally
		json.Unmarshal(byteValue, &config)

		defer jsonFileDat.Close()
		return config
	}
	// in case the env var is no defined just use the default config
	return readwebDefaultConfig()
}

// MAIN CONFIG FUNCS

// reads the hardcoded default config file
func readDefaultConfig() mainConfig {
	// create new config object of type Config
	var config mainConfig

	// open default file and log if error
	jsonFileDat, err := os.Open("/usr/local/gateTool/gateToolWeb/gateToolConfig.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("opened default file\n")
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
		fmt.Printf("opened file from env\n")
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
