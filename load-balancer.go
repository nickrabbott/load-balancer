package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("got / request")
	res, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Failed to dump the http request:", err)
		return
	}
	log.Println(string(res))
	io.WriteString(w, "load balancer\n")
	w.WriteHeader(200)
}

func main() {
	serverPort := os.Getenv("SERVER_PORT")

	http.HandleFunc("/", home)

	log.Printf("Server Port: %s\n", serverPort)

	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
    if errors.Is(err, http.ErrServerClosed) {
            log.Println("server closed")
    } else if err != nil {
            log.Printf("error starting server: %s\n", err)
            os.Exit(1)
    }
}
