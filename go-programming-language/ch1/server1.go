package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path=%s\n", r.URL.Path)
}

func main() {
	listen := "localhost:3000"

	http.HandleFunc("/", handler)

	fmt.Fprintf(os.Stdout, "Listening on '%s'\n", listen)

	err := http.ListenAndServe(listen, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
