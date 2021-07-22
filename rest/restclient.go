package rest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var API_KEY string
var BASE_URL string

// Essentially our contructor
func New(url string, apiKey string) {
	API_KEY = apiKey
	BASE_URL = url
}

func Get(route string) (string, error) {
	if BASE_URL == "" {
		log.Fatal("You must specify a base url to execute GET requests!")
		err := errors.New("You must specify a base url to execute GET requests!")
		return "", err
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest("GET", BASE_URL+route, nil)

	if err != nil {
		log.Fatal("Failed to create GET request. Error: ", err.Error())
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Api-Key", API_KEY)

	resp, err := client.Do(req)

	if err != nil {
		log.Error("Failed to execute GET request. Error: ", err.Error())
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("Your API Key is invalid or has expired")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error("Failed to read response body. Error: ", err.Error())
		return "", err

	}

	return string(body), nil

}
