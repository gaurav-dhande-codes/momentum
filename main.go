package main

import (
	"fmt"
	"rsc.io/quote"
	"net/http"
)

func main(){
	fmt.Println(quote.Go())
	
	// Create a basic HTTP request handler and register it to receives all incoming API requests to path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello!, you have requested: %s\n", r.URL.Path)
	})

	// Create a request handler to serve a Go proverb
	http.HandleFunc("/getGoProverb", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Your requested Go proverb is: %s\n", quote.Go())
	})

	// Start an HTTP server and listen for requests on port 80
	fmt.Println("Server Running...")
	http.ListenAndServe(":80", nil)
}