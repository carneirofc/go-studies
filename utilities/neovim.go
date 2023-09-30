package neovim

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Assets struct {
	ContentType        string `json:"content_type"`
	Size               int32  `json:"size"`
	Name               string `json:"name"`
	Url                string `json:"url"`
	BrowserDownloadUrl string `json:"browser_download_url"`
}

type Release struct {
	TagName      string    `json:"tag_name"`
	Url          string    `json:"url"`
	HtmlUrl      string    `json:"html_url"`
	AssetsUrl    string    `json:"assets_url"`
	Id           int32     `json:"id"`
	IsPreRelease bool      `json:"prerelease"`
	CreatedAt    time.Time `json:"created_at"`
	PublishedAt  time.Time `json:"published_at"`
	Assets       []Assets  `json:"assets"`
}

func gitApiBaseUrl() string {
	url := os.Getenv("GIT_API_URL")
	if url != "" {
		return url
	}
	return "https://api.github.com"
}

func getHttpClient() http.Client {
	client := http.Client{}
	return client
}

func DownloadTag() {
}

func ParseGitReleases(data []byte) ([]Release, error) {
	var releases []Release
	err := json.Unmarshal(data, &releases)
	if err != nil {
		return nil, err
	}
	return releases, nil
}

func GetReleaseByTag(tag string) (*Release, error) {
	if strings.TrimSpace(tag) == "" {
		return nil, fmt.Errorf("tag cannot be empty or whitespaces")
	}
	log.Fatal("not impplemented!")
	return nil, nil
}

func ListReleases(count int) ([]Release, error) {
	if count < 0 || count > 100 {
		return nil, fmt.Errorf("count must be between 1 and 100\n")
	}
	client := getHttpClient()
	req, err := http.NewRequest("GET", gitApiBaseUrl()+"/repos/neovim/neovim/releases", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	query := url.Values{}
	query.Add("per_page", fmt.Sprint(count))
	req.URL.RawQuery = query.Encode()
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request to '%s' return with status code '%s'\n", gitApiBaseUrl(), res.Status)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	releases, err := ParseGitReleases(data)
	if err != nil {
		return nil, err
	}
	return releases, nil
}
