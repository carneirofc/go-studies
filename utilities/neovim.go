package neovim

import (
	"encoding/json"
	"fmt"
	"io"
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

func Get(urlStr string, headers map[string]string, params map[string]string) ([]byte, error) {
	var err error
	client := getHttpClient()
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	query := url.Values{}
	for k, v := range params {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request '%s' status '%s'\n", gitApiBaseUrl(), res.Status)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return data, nil
}

func ParseGitReleases(data []byte) ([]Release, error) {
	var releases []Release
	err := json.Unmarshal(data, &releases)
	if err != nil {
		return nil, err
	}
	return releases, nil
}

/*
*  /repos/{owner}/{repo}/releases/tags/{tag}
 */
func GetReleaseByTag(tag string) (*Release, error) {
	if strings.TrimSpace(tag) == "" {
		return nil, fmt.Errorf("tag cannot be empty or whitespaces")
	}
	url := fmt.Sprintf("%s/repos/%s/%s/releases/tags/%s", gitApiBaseUrl(), "neovim", "neovim", tag)

	headers := make(map[string]string)
	headers["Accept"] = "application/vnd.github+json"
	headers["X-GitHub-Api-Version"] = "2022-11-28"

	projection := func(data []byte) (*Release, error) {
		var release Release
		err := json.Unmarshal(data, &release)
		if err != nil {
			return nil, err
		}
		return &release, nil
	}

	data, err := Get(url, headers, nil)
	if err != nil {
		return nil, fmt.Errorf("failed http request %s:%v", url, err)
	}
	parsed, err := projection(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse http request %s:%v", url, err)
	}
	return parsed, nil

}

func ListReleases(count int) ([]Release, error) {
	if count < 0 || count > 100 {
		return nil, fmt.Errorf("count must be between 1 and 100\n")
	}

	url := gitApiBaseUrl() + "/repos/neovim/neovim/releases"
	headers := make(map[string]string)
	headers["Accept"] = "application/vnd.github+json"
	headers["X-GitHub-Api-Version"] = "2022-11-28"

	params := make(map[string]string)
	params["per_page"] = fmt.Sprint(count)

	data, err := Get(url, headers, params)
	if err != nil {
		return nil, err
	}

	parsed, err := ParseGitReleases(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse http request %s:%v", url, err)
	}
	return parsed, nil
}
