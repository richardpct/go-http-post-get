// Spin up a webserver and displays the hostname
// If a POST request is performed then it writes the data in /tmp/post.txt file
// If a GET request is performed then it displays the /tmp/post.txt content
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	listen   = "0.0.0.0"
	port     = "8080"
	dataPost = "/tmp/post.txt"
)

func getContentFile(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		content = []byte("empty")
	}
	contentString := fmt.Sprintf("%s", content)
	return contentString
}

func writeContentFile(file string, content []byte) {
	contentByte := []byte(content)
	err := ioutil.WriteFile(file, contentByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	if r.Method == "GET" {
		content := getContentFile(dataPost)
		fmt.Fprintf(w, "The hostname is: "+hostname+
			"\nThe content is:\n"+content+"\n")
	} else if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		writeContentFile(dataPost, body)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listen+":"+port, nil))
}
