package utilities

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Download contents from the url into the dest
func DownloadFile(client *http.Client, urlStr string, dest io.Writer) (int64, error) {
	resp, err := client.Get(urlStr)
	if err != nil {
		return 0, fmt.Errorf("http get failed %s:%v", urlStr, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("http get failed %s %s", urlStr, resp.Status)
	}
	n, err := io.Copy(dest, resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to copy to destination:%v", err)
	}

	return n, nil
}

func CalcSha(fp io.Reader) ([]byte, error) {
	sha := sha256.New()
	rBuffer := make([]byte, 512)
	reader := bufio.NewReader(fp)
	for {
		n, err := reader.Read(rBuffer)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if n > 0 {
			_, werr := sha.Write(rBuffer[:n])
			if werr != nil {
				return nil, fmt.Errorf("failed to load buffer into sha256:%v", werr)
			}
		}

		if err == io.EOF {
			break
		}
	}
	return sha.Sum(nil), nil
}
func Get(urlStr string, headers map[string]string, params map[string]string) ([]byte, error) {
	return GetWithClient(http.DefaultClient, urlStr, headers, params)
}
func GetWithClient(client *http.Client, urlStr string, headers map[string]string, params map[string]string) ([]byte, error) {
	var err error
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
		return nil, fmt.Errorf("request '%s' status '%s'\n", urlStr, res.Status)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return data, nil
}
