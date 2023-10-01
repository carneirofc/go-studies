package neovim

import "time"

type Asset struct {
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
	Assets       []Asset   `json:"assets"`
}
