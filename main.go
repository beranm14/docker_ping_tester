package main

import (
	"crypto/tls"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	call_url := "https://example.com"
	if ru := os.Getenv("REMOTE_URL"); ru != "" {
		call_url = ru
	}
	app_port := "3000"
	if ap := os.Getenv("PORT"); ap != "" {
		app_port = ap
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello 🐱 from %q\n", html.EscapeString(r.URL.Path))
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		res, err := client.Get(call_url)
		if err != nil {
			log.Fatal(err)
		}
		data, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		fmt.Fprintf(w, "Called %s, got %d bytes:\n\n", call_url, len(data))
		fmt.Fprintf(w, "%s", data)
	})

	log.Println("Listening on port", app_port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", app_port), nil))
}
