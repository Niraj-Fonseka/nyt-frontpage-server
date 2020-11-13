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
	pwd, err := os.Getwd() //get current working directory
	if err != nil {
		log.Fatal(err)
	}
	staticPath := fmt.Sprintf("%s/static", pwd) //generate the path for the static files

	go fetchNYT(pwd) //fetch the static front page as a background job

	port := "80"

	http.Handle("/", http.FileServer(http.Dir(staticPath)))

	log.Printf("starting server on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

/*
fetchNYT
- runs as a background job every 1 hour and checks if a new page is available
*/
func fetchNYT(pwd string) {

	path := fmt.Sprintf("%s/%s", pwd, "static/nyt.pdf")
	for {
		log.Println("Starting fetch ....")
		resp, err := http.Get(generateURL())

		if resp.StatusCode != 200 { //if the next day is not available will get a 404
			time.Sleep(1 * time.Hour) //run every hour
			continue
		}

		if err != nil {
			log.Println(err)
			continue
		}

		defer resp.Body.Close()

		// Create the new pdf file
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
		time.Sleep(1 * time.Hour) //run every hour
	}
}

/*
generateURL
- get the current date and generate the url
*/

func generateURL() string {
	return fmt.Sprintf("https://static01.nyt.com/images/%s/nytfrontpage/scan.pdf", time.Now().Format("2006/01/02"))
}
