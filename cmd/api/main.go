package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", helloHandler)
	var port = os.Getenv("APP_PORT")
	log.Println("Listing for requests at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
