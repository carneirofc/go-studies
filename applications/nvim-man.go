package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	neovim "github.com/carneirofc/go-studies/neovim"
)

func main() {
	data, err := neovim.GetReleaseByTag("nightly")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data.HtmlUrl)

	if runtime.GOARCH != "amd64" {
		log.Fatalf("unsupported architecture %s\n", runtime.GOARCH)
	}

	isValidAsset := func(v neovim.Asset) bool {
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
		if !isValidAsset(v) {
			continue
		}

		fmt.Println(v.BrowserDownloadUrl)
	}
	// TODO: Impplement download and checksum validation if exists

}
