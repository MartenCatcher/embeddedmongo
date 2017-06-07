package embeddedmongo

import (
	"archive/zip"
	"io"
	"os"
	"regexp"
)

func Extract(d *Distribution, command Command) (string, error) {
	workDir := GetWorkDir(d)
	path := workDir + GetDistributionName(d)
	r, err := zip.OpenReader(path)
	if err != nil {
		return "", err
	}
	defer r.Close()

	regexFilename, err := regexp.Compile(".*/([^/]*)$")
	if err != nil {
		return "", err
	}

	//TODO: use command for file choosing
	regexCommand, err := regexp.Compile("mongod(.exe)?$")
	if err != nil {
		return "", err
	}

	tmp := GetTmpDir(d)
	CreateDir(tmp)
	if err != nil {
		return "", err
	}
	app := func() string {
		for _, file := range r.File {
			name := regexFilename.ReplaceAllString(file.Name, "$1")
			isExec := regexCommand.MatchString(name)
			//log.Printf("%v\t\t\t%v\n", file.Name, name)
			if isExec {
				extractOneFile(file, tmp+name)
				return name
			}
		}
		return ""
	}()

	return app, nil
}

func extractOneFile(file *zip.File, path string) {
	reader, _ := file.Open()
	//TODO: catch error
	defer reader.Close()

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeTemporary)
	defer f.Close()
	//TODO: catch error
	io.Copy(f, reader)
}
