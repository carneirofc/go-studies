package main

import (
	"flag"
	"log"
)

type t_Config struct {
	Tag       string
	ListCount int
	List      bool
	Get       bool
}

var config t_Config

func init() {
	flag.StringVar(&config.Tag, "tag", "nightly", "selected tag to download")
	flag.BoolVar(&config.List, "list", false, "list releases")
	flag.BoolVar(&config.Get, "get", false, "get releases")
	flag.IntVar(&config.ListCount, "list-count", 2, "list releases count")
	flag.Parse()

	if config.ListCount < 1 || config.ListCount >= 100 {
		flag.Usage()
		log.Fatal("config must be >= 1 < 100")
	}
}

func main() {
	if config.List {
		list()
		return
	}
	if config.Get {
		get()
		return
	}
}
