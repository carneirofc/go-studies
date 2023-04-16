package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	channel := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(channel, url)
	}

	for range os.Args[1:] {
		fmt.Printf("%s\n", <-channel)
	}
	fmt.Printf("Elapsed %.2f\n", time.Since(start).Seconds())
}

func fetch(channel chan<- string, url string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprintf("%s %v", url, err)
		return
	}

	count, err := io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	if err != nil {
		channel <- fmt.Sprintf("%s %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	channel <- fmt.Sprintf("%0.2fs: %7d %s", secs, count, url)
}
