package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func get(_url string) error {
	url := _url
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("HTTP request returned with a status code '%d'", res.StatusCode)
	}

	// buf, err := ioutil.ReadAll(res.Body)
	_, err = io.Copy(os.Stdout, res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	// fmt.Printf("%s", buf)
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) < 0 {
		fmt.Fprintf(os.Stderr, "Insufficient arguments\n")
		os.Exit(-1)
	}

	for _, url := range args {

		err := get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
}
