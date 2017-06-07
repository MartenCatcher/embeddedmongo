package embeddedmongo

import (
	"io"
	"net/http"
	"os"
)

func Download(fileName string, workDir string, url string) (err error) {
	filePath := workDir + fileName
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		CreateDir(workDir)

		out, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer out.Close()

		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}

		return nil
	}
	return nil
}
