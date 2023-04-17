package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %q %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] %q\n", k, v)
	}
	body := r.Body
	if body != nil {
		data, err := io.ReadAll(body)
		if err != nil {
			log.Print(err)
		}
		fmt.Fprintf(w, "Body = %q\n", string(data))
	}
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
