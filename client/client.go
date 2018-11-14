package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gustavotero7/go-conekta/models"
)

const (
	// APIURL Base API url
	APIURL = "https://api.conekta.io"
	// APIVersion Current API version
	APIVersion = "2.0.0"
)

// APIKey Conecta secret key
var APIKey = ""

// Client conecta api client
// We use this struct to mock the client for unit testing
type client struct {
	apiURL     string
	apiVersion string
	apiKey     string
}

var c *client

func getClient() *client {
	if c == nil {
		c = &client{
			apiURL:     APIURL,
			apiVersion: APIVersion,
			apiKey:     APIKey,
		}
	}
	return c
}

// Get do get request
func Get(path string) ([]byte, error) {
	return getClient().send(http.MethodGet, path, nil)
}

// Post do post request
func Post(path string, payload interface{}) ([]byte, error) {
	return getClient().send(http.MethodPost, path, payload)
}

// Put do put request
func Put(path string, payload interface{}) ([]byte, error) {
	return getClient().send(http.MethodPut, path, payload)
}

// Delete do delete request
func Delete(path string) ([]byte, error) {
	return getClient().send(http.MethodDelete, path, nil)
}

func (c *client) send(method, path string, v interface{}) ([]byte, error) {
	payload, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.apiURL, path), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", fmt.Sprintf("application/vnd.conekta-v%s+json", c.apiVersion))
	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(c.apiKey, "")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// If status isn't OK (200) then try to decode it into a conekta error
	// If decode fails then return decode (Unmarshal) error and nil response
	if res.StatusCode != 200 {
		conektaError := &models.ConektaError{}
		if err := json.Unmarshal(body, conektaError); err != nil {
			return nil, err
		}
		return nil, conektaError
	}

	return body, nil
}

// SetClientURL re-sets api URL, bu default is "https://api.conekta.io"
// For now this is only used for testing purposes
func SetClientURL(apiURL string) {
	getClient().apiURL = apiURL
	log.Println("Client mocked")
}
