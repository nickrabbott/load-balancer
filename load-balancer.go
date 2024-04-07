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

var LAST_USED bool

func proxyGet(r *http.Request) (*http.Response, error) {
	var url string
	if LAST_USED {
		url = fmt.Sprintf("http://apache1%s", r.URL.Path)
	} else {
		url = fmt.Sprintf("http://apache2%s", r.URL.Path)
	}
	resp, err := http.Get(url)
	if err != nil {
		msg := "Failed to proxy GET request for %s"
		log.Printf(msg, url)
		return nil, err
	}
	LAST_USED = !LAST_USED // Server the other one next time
	return resp, nil
}

func proxyRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("got / request")
	res, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Failed to dump the http request:", err)
		return
	}
	log.Println(string(res))
	// io.WriteString(w, "load balancer\n")
	// w.WriteHeader(200)
	switch r.Method {
	case "GET":
		getResp, getErr := proxyGet(r)
		if getErr != nil {
			log.Println(getErr)
			w.WriteHeader(502)
			io.WriteString(w, getErr.Error())
		} else {
        	w.WriteHeader(getResp.StatusCode)
        	// Get the response body
        	body, err := io.ReadAll(getResp.Body)
        	if err != nil {
        	        msg := "Failed streaming response body"
        	        log.Println(msg)
        	        return
        	}
        	w.Write(body)
		}
	default:
		log.Printf("%v method not implemented", r.Method)
	}
}

func main() {
	serverPort := os.Getenv("SERVER_PORT")

	http.HandleFunc("/", proxyRequest)

	log.Printf("Server Port: %s\n", serverPort)

	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
    if errors.Is(err, http.ErrServerClosed) {
            log.Println("server closed")
    } else if err != nil {
            log.Printf("error starting server: %s\n", err)
            os.Exit(1)
    }
}
