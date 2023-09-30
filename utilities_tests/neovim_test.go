package test

import (
	"os"
	"testing"

	neovim "github.com/carneirofc/go-studies/utilities"
)

func Test_NeovimReleasesGet(t *testing.T) {
	releases, err := neovim.ListReleases(2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("successfully parsed '%d' releases\n", len(releases))
	for _, release := range releases {
		t.Logf("\n%s\n", release.TagName)
	}
}

func Test_GitReleases(t *testing.T) {
	filename := "mock.json"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	releases, err := neovim.ParseGitReleases(bytes)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("successfully parsed '%d' releases\n", len(releases))
	for _, release := range releases {
		t.Logf("\n%s\n", release.TagName)
		// s, _ := json.MarshalIndent(release, "", "\t")
		// t.Logf("\n%s\n", s)
	}
}
