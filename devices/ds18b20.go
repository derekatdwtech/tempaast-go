package devices

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var ErrReadSensor = errors.New("Failed to sensor temperature. Ensure you have connected your device to your RaspberryPi")

func SetupDevices() {
	log.Info("Initializing w1-gpio")
	_, gErr := exec.Command("modprobe w1-gpio").Output()

	if gErr != nil {
		log.Error("Failed to initialize w1-gpio. Message: ", gErr)
		os.Exit(1)
	}

	log.Info("Initializing w1-therm")

	_, tErr := exec.Command("modprobe w1-therm").Output()

	if tErr != nil {
		log.Error("Failed to initialize w1-therm. Message: ", tErr)
		os.Exit(1)
	}
}

func ReadTemperature(sensor string) float64 {
	data, err := ioutil.ReadFile(sensor)

	if err != nil {
		log.Error("Failed to access probe directory ", sensor)
		return 0.0
	}

	raw := string(data)

	if !strings.Contains(raw, " YES") {
		log.Warn("Waiting for sensor to come online")
		time.Sleep(10 * time.Second)
		ReadTemperature(sensor)
	}

	i := strings.LastIndex(raw, "t=")

	if i == -1 {
		log.Error("Failed to get index of probe file.")
		return 0.0
	}

	c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)

	if err != nil {
		log.Error("Failed to parse temperature reading from probe")
		return 0.0
	}

	temp := CelciusToFarenheit(c / 1000.0)

	log.Info("Read temperature as: ", temp)

	return temp
}
func CelciusToFarenheit(celcius float64) float64 {
	return celcius*9.0/5.0 + 32
}
