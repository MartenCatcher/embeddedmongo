package embeded_mongo

import (
	"fmt"
	"archive/zip"
	"regexp"
	"os"
	"io"
)

func Extract(d *Distribution, command Command) {
	workDir := fmt.Sprintf("%v/%v/", d.Dir, d.Os)
	path := fmt.Sprintf("%v%v-%v-%v-%v.%v", workDir, "mongodb", d.Os, d.Platform, d.Version, d.Extension)
	r, err := zip.OpenReader(path)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	regexFilename, err := regexp.Compile(".*/([^/]*)$")
	if err != nil {
		panic(err)
	}

	//TODO: use command for file choosing
	regexCommand, err := regexp.Compile("mongod(.exe)?$")
	if err != nil {
		panic(err)
	}

	app := func() string {
		for _, file := range r.File {
			name := regexFilename.ReplaceAllString(file.Name, "$1")
			isExec := regexCommand.MatchString(name)
			fmt.Printf("%v\t\t\t%v\n", file.Name, name)
			if isExec {
				path = workDir + name
				extractOneFile(file, path)
				return path
			}
		}
		return ""
	}()

	fmt.Printf("%+v\n", app)
}

func extractOneFile(file *zip.File, path string) {
	reader, _ := file.Open()
	//TODO: catch error
	defer reader.Close()

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeTemporary)
	//TODO: catch error
	io.Copy(f, reader)
}