package embeddedmongo

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Download the distribution or report an error
func Download(d *Distribution) (file string, err error) {

	workDir := GetWorkDir(d)
	filePath := workDir + GetDistributionName(d)

	if _, err := os.Stat(filePath); os.IsExist(err) {
		return filePath, errors.New("file exists")
	}

	url := GetDistributionUrl(d)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("%v (%v)", resp.Status, url))
	}

	err = CreateDir(workDir)
	if err != nil {
		return "", err
	}

	out, err := os.Create(filePath)
	if err != nil {
		os.RemoveAll(workDir)
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		os.RemoveAll(filePath)
		return "", err
	}

	return filePath, nil
}
