package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/carneirofc/go-studies/neovim"
	"github.com/carneirofc/go-studies/utilities"
)

func list(flags *tListFlags) {
	data, err := neovim.ListReleases(flags.Count)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range data {
		fmt.Printf("%s\t%s\n", d.HtmlUrl, d.PublishedAt)
		if !flags.All {
			continue
		}
		for _, asset := range d.Assets {
			if strings.Contains(asset.Name, "sha256sum") {
				continue
			}
			if strings.Contains(asset.Name, "zsync") {
				continue
			}

			fmt.Printf("%s\t%s\n", asset.BrowserDownloadUrl, utilities.FormatSizeBit(int64(asset.Size)))
		}
	}
}

func filename(flags *tInstallFlags) string {
	fname := "nvim"
	if flags.Format != ".appimage" {

		switch os := runtime.GOOS; os {
		case "windows":
			fname += "-win64"
		case "linux":

			fname += "-linux64"
		default:
			fname += "macos"
			return fname
		}
	}
	fname += flags.Format
	return fname
}

func install(flags *tInstallFlags) {
	stat, err := os.Stat(flags.Prefix)
	if os.IsNotExist(err) {
		err = os.Mkdir(flags.Prefix, 0755)
		if err != nil {
			log.Fatalf("could not create directory %s:%v", flags.Prefix, err)
		}
	}

	if !stat.IsDir() {
		log.Fatalf("%s is not a directory", flags.Prefix)
	}

	data, err := neovim.GetReleaseByTag(flags.Tag)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("listing content from %s\n", data.HtmlUrl)

	if runtime.GOARCH != "amd64" {
		log.Fatalf("unsupported architecture %s\n", runtime.GOARCH)
	}
	fname, err := filepath.Abs(filename(flags))
	if err != nil {
		log.Fatalf("failed to get abspath from %s:%v", fname, err)
	}
	// NOTE: Download sha from the interwebs!
	stat, err = os.Stat(fname)
	if err == nil {
		if flags.Force {
			log.Printf("removing file %s", fname)
			// TODO: Check sha256sum so we can detect if we have to upgrade or not!
			os.Remove(fname)
		} else {
			log.Fatalf("file %s already exists, use flag '--force' in order to overwrite it", fname)
		}
	}
	fp, err := os.Create(fname)
	if err != nil {
		log.Fatalf("failed to open file %s:%v", fname, err)
	}

	urlStr := fmt.Sprintf("%s/neovim/neovim/releases/download/%s/%s", neovim.GitBaseUrl(), flags.Tag, filepath.Base(fname))
	client := http.DefaultClient
	n, err := utilities.DownloadFile(client, urlStr, fp)
	if err != nil {
		log.Fatalf("failed to download %s:%v", urlStr, err)
	}
	log.Printf("downloaded %s %s\n", fname, utilities.FormatSizeBit(n))

	////assetFilterForPlatform := func(v neovim.Asset) bool {
	////	if !(runtime.GOOS == "windows") && !(runtime.GOOS == "linux") {
	////		return false
	////	}

	////	if (runtime.GOOS == "windows") && strings.Contains(v.BrowserDownloadUrl, "win") {
	////		return true
	////	}
	////	if (runtime.GOOS == "linux") && strings.Contains(v.BrowserDownloadUrl, "linux") {
	////		return true
	////	}
	////	if (runtime.GOOS == "linux") && strings.Contains(v.BrowserDownloadUrl, "appimage") {
	////		return true
	////	}

	////	return false
	////}

	////for _, v := range data.Assets {
	////	if !assetFilterForPlatform(v) {
	////		continue
	////	}

	// //	// TODO: Impplement download and checksum validation if exists
	// //	log.Fatalln("TODO: Impplement download and checksum validation if exists")
	// //	fmt.Println(v.BrowserDownloadUrl)
	// //}
}
