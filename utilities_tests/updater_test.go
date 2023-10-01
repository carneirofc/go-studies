package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/carneirofc/go-studies/utilities"
)

func Test_Formatting(t *testing.T) {
	t.Log(utilities.FormatSize(int64(1024 * 100)))
	t.Log(utilities.FormatSizeBit(int64(1024 * 100)))
}

func Test_DownloadFile(t *testing.T) {
	t.Fatal()
	if _, err := os.Stat("./tmp"); os.IsNotExist(err) {
		err := os.Mkdir("./tmp", 0755)
		if err != nil {
			t.Errorf("failed to create directory:%v", err)
		}
	}
	defer os.RemoveAll("./tmp/")

	fd, err := os.CreateTemp("./tmp", "nvim.appimage.sha256sum")
	if err != nil {
		t.Error(err)
	}

	urlStr := "https://github.com/neovim/neovim/releases/download/nightly/nvim.appimage.sha256sum"
	t.Logf("downloading contents into %s", fd.Name())

	n, err := utilities.DownloadFile(http.DefaultClient, urlStr, fd)
	if err != nil {
		t.Errorf("failed to download file from %s:%v", urlStr, err)
	}
	t.Logf("downloaded %s", utilities.FormatSizeBit(n))
	fd.Close()

	// Downloaing executable
	fd, err = os.CreateTemp("./tmp", "nvim.appimage")
	if err != nil {
		t.Error(err)
	}

	urlStr = "https://github.com/neovim/neovim/releases/download/nightly/nvim.appimage"
	t.Logf("downloading contents into %s", fd.Name())

	n, err = utilities.DownloadFile(http.DefaultClient, urlStr, fd)
	if err != nil {
		t.Errorf("failed to download file from %s:%v", urlStr, err)
	}
	t.Logf("downloaded %s", utilities.FormatSizeBit(n))
}

func Test_ExistsInPath(t *testing.T) {
	name := "nvim"
	instances, err := utilities.FindInPath(name)
	if err != nil {
		t.Fatal(err)
	}
	for _, instance := range instances {
		t.Log(instance)
	}
}

func Test_CalcSystemSha(t *testing.T) {
	// "https://github.com/neovim/neovim/releases/download/nightly/nvim.appimage"
	// "https://github.com/neovim/neovim/releases/download/nightly/nvim.appimage.sha256sum"
	instances, err := utilities.FindInPath("nvim")
	if err != nil {
		t.Fatal(err)
	}
	for _, fpath := range instances {
		stat, err := os.Stat(fpath)
		if err != nil {
			t.Errorf("path %s : %v", fpath, err)
		}
		fp, err := os.Open(fpath)
		if err != nil {
			t.Errorf("path %s : %v", fpath, err)
		}

		sha, err := utilities.CalcSha(fp)
		if err != nil {
			t.Fatalf("failed to calc sha %v", err)
		}
		fp.Close()

		t.Logf("%s\t%s\t%x", fpath, utilities.FormatSizeBit(stat.Size()), sha)
	}
}

func Test_CalcSha(t *testing.T) {
	fpath := "./nvim.appimage"
	fpathSha := "./nvim.appimage.sha256sum"
	xdata, err := ioutil.ReadFile(fpathSha)
	if err != nil {
		t.Error(err)
	}
	idx := bytes.IndexFunc(xdata, func(r rune) bool { return string(r) == " " })
	xsha := string(xdata[:idx])

	t.Logf("expected %s", xsha)

	fp, err := os.Open(fpath)
	if err != nil {
		t.Errorf("path %s:%v", fpath, err)
	}
	defer fp.Close()

	shab, err := utilities.CalcSha(fp)
	if err != nil {
		t.Error(err)
	}
	sha := fmt.Sprintf("%x", shab)
	if xsha != sha {
		t.Errorf("expected sha %s differs from %s", xsha, sha)
	}
	t.Logf("%s\t%s", fpath, sha)
}
