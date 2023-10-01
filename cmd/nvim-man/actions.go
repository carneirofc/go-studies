package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/carneirofc/go-studies/neovim"
)

func list() {
	data, err := neovim.ListReleases(config.ListCount)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range data {
		fmt.Printf("%s\t%s\n", d.HtmlUrl, d.PublishedAt)
	}
}

func get() {
	data, err := neovim.GetReleaseByTag(config.Tag)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("listing content from %s\n", data.HtmlUrl)

	if runtime.GOARCH != "amd64" {
		log.Fatalf("unsupported architecture %s\n", runtime.GOARCH)
	}

	assetFilterForPlatform := func(v neovim.Asset) bool {
		if !(runtime.GOOS == "windows") && !(runtime.GOOS == "linux") {
			return false
		}

		if (runtime.GOOS == "windows") && strings.Contains(v.BrowserDownloadUrl, "win") {
			return true
		}
		if (runtime.GOOS == "linux") && strings.Contains(v.BrowserDownloadUrl, "linux") {
			return true
		}
		if (runtime.GOOS == "linux") && strings.Contains(v.BrowserDownloadUrl, "appimage") {
			return true
		}

		return false
	}

	for _, v := range data.Assets {
		if !assetFilterForPlatform(v) {
			continue
		}

		fmt.Println(v.BrowserDownloadUrl)
	}
	// TODO: Impplement download and checksum validation if exists
	log.Fatalln("TODO: Impplement download and checksum validation if exists")
}
