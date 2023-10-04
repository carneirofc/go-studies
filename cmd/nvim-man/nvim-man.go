package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

type tInstallFlags struct {
	Tag    string
	Prefix string
	Format string
	Force  bool
}
type tListFlags struct {
	Count int
	All   bool
}

func init() {
	log.SetFlags(0)
}
func main() {
	help := func() {
		fmt.Fprintf(os.Stderr, "valid commands: \"list\", \"install\"\n")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		help()
	}

	switch cmd := os.Args[1]; cmd {
	case "list":
		config := tListFlags{}
		fList := flag.NewFlagSet("list", flag.ExitOnError)
		fList.IntVar(&config.Count, "count", 2, "list releases count")
		fList.BoolVar(&config.All, "all", false, "list all artifacts")
		fList.Parse(os.Args[2:])
		list(&config)

	case "install":
		config := tInstallFlags{}
		fInstall := flag.NewFlagSet("install", flag.ExitOnError)
		fInstall.BoolVar(&config.Force, "force", false, "overwrite existing files")
		fInstall.StringVar(&config.Tag, "tag", "nightly", "selected tag to download")
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("could not get current working directory %v", err)
		}
		fInstall.StringVar(&config.Prefix, "prefix", wd, "destination to install, defaults to cwd")

		var defaultFormat string
		if runtime.GOOS == "linux" {
			defaultFormat = ".appimage"
		} else if runtime.GOOS == "windows" {
			defaultFormat = ".zip"
		} else {
			defaultFormat = ".tar.gz"
		}
		fInstall.StringVar(&config.Format, "format", defaultFormat, "format to download [.tar.gz|.appimage|.zip|.msi]")
		fInstall.Parse(os.Args[2:])
		install(&config)

	default:
		help()
	}
}
