package main

import (
	"flag"
	"fmt"
	"os"
	"tempaast/devices"
	"time"
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

	devices.SetupDevices()

	for true {
		devices.ReadTemperature(probeDir)
		time.Sleep(10 * time.Second)
	}
}
