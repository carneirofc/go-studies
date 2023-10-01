package utilities

import (
	"fmt"
	"math"
	"os"
	"path"
	"runtime"
	"strings"
)

func FindInPath(fname string) ([]string, error) {
	paths := os.Getenv("PATH")
	isWindows := runtime.GOOS == "windows"
	isLinux := runtime.GOOS == "linux"
	var separator string
	if isWindows {
		separator = ";"
	}
	if isLinux {
		separator = ":"
	}
	if separator == "" {
		return nil, fmt.Errorf("invalid separator '%s'", separator)
	}

	matches := []string{}
	for _, p := range strings.Split(paths, separator) {
		fpath := path.Join(p, "nvim")
		stat, err := os.Stat(fpath)
		if os.IsNotExist(err) {
			continue
		}
		if stat.IsDir() {
			continue
		}
		matches = append(matches, fpath)
	}
	return matches, nil
}
func formatSize(size int64, list []string, step int) string {
	if size < int64(step) {
		return fmt.Sprintf("%d%s", size, list[0])
	}
	for i, lbl := range list {

		lowerBound := math.Floor(math.Pow(float64(step), float64(i)))
		upperBound := math.Floor(math.Pow(float64(step), float64(i+1)))
		if i == len(list)-1 {
			upperBound = -1
		}

		if float64(size) >= lowerBound && float64(size) < upperBound {
			return fmt.Sprintf("%0.3f%s", float64(size)/math.Pow(float64(step), float64(i)), lbl)
		}
	}
	return fmt.Sprintf("%d", size)
}
func FormatSize(size int64) string {
	list := []string{"B", "KB", "MB", "GB"}
	step := 1000
	return formatSize(size, list, step)
}
func FormatSizeBit(size int64) string {
	list := []string{"B", "KiB", "MiB", "GiB"}
	step := 1024
	return formatSize(size, list, step)
}

// t.Logf("%s\t%d MB\t%x", fpath, stat.Size()/1024/1024, sha)
