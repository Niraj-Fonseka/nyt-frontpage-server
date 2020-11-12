package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	go fetchNYT()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}

func fetchNYT() {
	for {
		log.Println("Starting fetch ....")

		resp, err := http.Get("https://static01.nyt.com/images/2020/11/11/nytfrontpage/scan.pdf")
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()

		// Create the file
		out, err := os.Create("./static/scan.pdf")
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
