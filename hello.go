package main

import "fmt"
import "log"
import "net/http"
import "golang.org/x/net/html"

func main() {
	fmt.Println("Starting server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hello, world. we are at %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}