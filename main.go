package main

import (
	"fmt"
	"os"
	"time"
	"net/http"
	"log"
)

// https://stackoverflow.com/a/45909815/385205
func wrapHandlerWithLogging(wrappedHandler http.Handler) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
       fmt.Printf("--> %s %s\n", req.Method, req.URL.Path)
       wrappedHandler.ServeHTTP(w, req)
    })
}

// https://gobyexample.com/http-servers
func hello(w http.ResponseWriter, req *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	log.Println("Request recieved")

	fmt.Fprintf(w, "hostname: " + name + "\n")

	t := time.Now()
	fmt.Fprintf(w, "timenow: " + t.String() )
	
	fmt.Println("timenow: " + t.String() )
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	fmt.Fprintf(w, "Date: %v\n", time.Now().Format(time.RFC1123))

}	

func main() {
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/", wrapHandlerWithLogging(http.HandlerFunc(hello)))
	
	fmt.Println( "Starting application server in port 80")
	http.ListenAndServe(":80", nil)
}
