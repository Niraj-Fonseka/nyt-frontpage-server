package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	staticPath := fmt.Sprintf("%s/static", pwd)
	go fetchNYT(pwd)
	http.Handle("/", http.FileServer(http.Dir(staticPath)))
	http.ListenAndServe(":3000", nil)
}

func fetchNYT(pwd string) {

	path := fmt.Sprintf("%s/%s", pwd, "static/nyt.pdf")

	log.Println(path)
	for {
		log.Println("Starting fetch ....")

		resp, err := http.Get("https://static01.nyt.com/images/2020/11/11/nytfrontpage/scan.pdf")
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()

		// Create the file
		out, err := os.Create(path)
		if err != nil {
			log.Println(err)
			continue
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Done fetch ....")
		time.Sleep(60 * time.Second)
	}
}
