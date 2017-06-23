package embeddedmongo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
)

// Extract the distribution and return the directory files list
func Extract(d *Distribution) ([]string, error) {

	files := []string{}

	path := GetWorkDir(d) + GetDistributionName(d)
	temp := GetTmpDir(d)

	err := CreateDir(temp)
	if err != nil {
		log.Println(err)
		return files, err
	}

	switch d.Extension {
	case "zip":
		err = archiver.Zip.Open(path, temp)
	case "tgz":
		err = archiver.TarGz.Open(path, temp)
	default:
		return files, errors.New(fmt.Sprintf("not supported archive: %v", d.Extension))
	}

	err = filepath.Walk(temp, func(fpath string, f os.FileInfo, err error) error {
		files = append(files, fpath)
		return nil
	})

	return files, nil
}
