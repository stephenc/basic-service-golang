package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)
	listenAddr := ":8080"
	if l, ok := os.LookupEnv("LISTEN_ADDR"); ok {
		listenAddr = l
	}

	http.ListenAndServe(listenAddr, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	greeting := os.Getenv("CUSTOM_GREETING")
	if greeting == "" {
		greeting = "Hello"
	}
	name := r.URL.Path[1:]
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "%s, %s!", greeting, name)
}
