package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex
var counter int64 = 0

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	mu.Unlock()

	fmt.Fprintf(w, "URL.path=%s\n", r.URL.Path)
}
func count(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	mu.Unlock()

	fmt.Fprintf(w, "Count=%d\n", counter)
}

func main() {
	listen := "localhost:3000"
	http.HandleFunc("/count", count)
	http.HandleFunc("/", handler)

	fmt.Fprintf(os.Stdout, "Listening on '%s'\n", listen)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
