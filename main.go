package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/prometheus/alertmanager/template"
)

const token = "guBsWD3wzEsk9c9O0txnHeDSHz9zBz76xSfmTcF40Gb"
const line_uri = "https://notify-api.line.me/api/notify"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("--require port--")
	}

	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/hook", hookHandler)

	log.Printf("server start using port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	d := template.Data{}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lineChannel(d)
	log.Println("------")
}

func lineChannel(info template.Data) {
	bearer := "Bearer " + token

	c := &http.Client{}
	body := url.Values{}

	body.Set("message", "test message")

	r, _ := http.NewRequest("POST", line_uri, bytes.NewBufferString(body.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", bearer)

	res, err := c.Do(r)
	if err != nil {
		log.Printf("request error %v", err)
	}
	defer res.Body.Close()

	f, _ := ioutil.ReadAll(res.Body)

	log.Printf("result: %v", string(f))
}
