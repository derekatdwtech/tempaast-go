package main

import (
	"flag"
	"fmt"
	"os"
	"tempaast/devices"
	"tempaast/rest"
	"time"

	log "github.com/sirupsen/logrus"
)

var probeDir string
var probeName string
var apiKey string

func init() {
	const (
		probeDirUsage  = "Specify the path of your device."
		probeNameUsage = "Provide the unique nickname for your device."
		apiKeyUsage    = "Provide your tempaast API Key"
	)

	flag.StringVar(&probeDir, "d", "", probeDirUsage)
	flag.StringVar(&probeName, "n", "", probeNameUsage)
	flag.StringVar(&apiKey, "k", "", apiKeyUsage)
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Initialize REST Client
	rest.New("https://meatmonitorapi.azurewebsites.net/", apiKey)
	// Validate API Key
	_, err := rest.Get("api/key/validate")
	if err != nil {
		log.Fatal("Your API Key is invalid or has expired.")
		os.Exit(77)
	}

	devices.SetupDS18B20()

	for true {
		devices.ReadDS18B20(probeDir)
		time.Sleep(10 * time.Second)
	}
}
