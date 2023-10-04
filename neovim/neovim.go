package neovim

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/carneirofc/go-studies/utilities"
)

func GitBaseUrl() string {
	url := os.Getenv("GIT_BASE_URL")
	if url != "" {
		return url
	}
	return "https://github.com"
}
func GitApiBaseUrl() string {
	url := os.Getenv("GIT_API_URL")
	if url != "" {
		return url
	}
	return "https://api.github.com"
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
	url := fmt.Sprintf("%s/repos/%s/%s/releases/tags/%s", GitApiBaseUrl(), "neovim", "neovim", tag)

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

	data, err := utilities.Get(url, headers, nil)
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

	url := GitApiBaseUrl() + "/repos/neovim/neovim/releases"
	headers := make(map[string]string)
	headers["Accept"] = "application/vnd.github+json"
	headers["X-GitHub-Api-Version"] = "2022-11-28"

	params := make(map[string]string)
	params["per_page"] = fmt.Sprint(count)

	data, err := utilities.Get(url, headers, params)
	if err != nil {
		return nil, err
	}

	parsed, err := ParseGitReleases(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse http request %s:%v", url, err)
	}
	return parsed, nil
}
