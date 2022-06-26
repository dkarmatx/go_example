package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Config struct {
	Host string `json:"Host"`
	Tag  string `json:"Tag"`
}

func main() {
	log.Print("Loading config ...")

	var buff []byte
	var cfg Config

	if readed, err := ioutil.ReadFile("config.json"); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	} else {
		buff = readed
	}
	if err := json.Unmarshal(buff, &cfg); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		respBody := fmt.Sprintf(
			`<html><body>
				<h3> *** WebApp3000 *** </h3>
				<p>My tag is: <b>%s</b></p>
			</body></html>`, cfg.Tag,
		)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(respBody))
	})

	log.Print("Application is starting ...")
	log.Fatal(http.ListenAndServe(cfg.Host, nil))
}
