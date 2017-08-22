package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func getObsfucate(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("url")
	obsfURL := base64.RawURLEncoding.EncodeToString([]byte(url))
	fmt.Fprint(w, "http://"+req.Host+"/?u="+obsfURL)
}

func getRedirect(w http.ResponseWriter, req *http.Request) {
	obsfURL := req.URL.Query().Get("u")
	decodedURL, err := base64.RawURLEncoding.DecodeString(obsfURL)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintln(w, "ERROR: Decoding failed!")
		return
	}
	http.Redirect(w, req, string(decodedURL), http.StatusTemporaryRedirect)
}

func main() {
	http.Handle("/", http.HandlerFunc(getRedirect))
	http.Handle("/obsfucate", http.HandlerFunc(getObsfucate))
	http.ListenAndServe(":8080", nil)
}
