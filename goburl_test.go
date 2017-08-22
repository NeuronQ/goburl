package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	baseURL               = "http://localhost:8080"
	urlToObsfucate        = "https://golang.org/pkg/encoding/base64/"
	expectedObsfucatedURL = baseURL + "/?u=aHR0cHM6Ly9nb2xhbmcub3JnL3BrZy9lbmNvZGluZy9iYXNlNjQv"
)

func TestHealthCheckHandler(t *testing.T) {
	obsfucateRequestURL := baseURL + "/obsfucate?url=" + urlToObsfucate
	resp, err := http.Get(obsfucateRequestURL)
	if err != nil {
		fmt.Printf("error performing obsfucation request %s: %v\n", obsfucateRequestURL, err)
		t.Fail()
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading obsfucation request response's body: %v\n", err)
		t.Fail()
		return
	}
	obsfucatedURL := string(body)

	// test actual response
	if obsfucatedURL != expectedObsfucatedURL {
		fmt.Printf("obsfucated url %s does not match expected %s\n", obsfucatedURL, expectedObsfucatedURL)
		t.Fail()
	}

	// test redirect
	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}
	resp, err = client.Get(obsfucatedURL)
	if err != nil {
		if strings.Contains(err.Error(), "Redirect") {
			if resp.StatusCode != http.StatusTemporaryRedirect {
				fmt.Printf("not the expected redirect status code on request to obsfucated url %v: %v but expected %v\n",
					obsfucatedURL, resp.StatusCode, http.StatusTemporaryRedirect)
				t.Fail()
			}
			if resp.Header["Location"][0] != urlToObsfucate {
				fmt.Printf("redirect to unexpected location %s instead of %s\n", resp.Header["Location"][0], urlToObsfucate)
			}
		} else {
			fmt.Printf("error when performing request to obsfucataed url: %v\n", err)
			t.Fail()
			return
		}
	} else {
		fmt.Printf("a redirect response was expected\n")
		t.Fail()
		return
	}
}
