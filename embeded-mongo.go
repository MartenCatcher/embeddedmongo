package embeded_mongo

import (
	"time"
	"archive/zip"
	"regexp"
	"fmt"
	"os"
	"io"
)

type (
	Command int
	Version string
)

const (
	Mongod Command = iota
)

const (
	V3_4_1 Version = "3.4.1"
)

type Distribution struct {
	Dir       string
	Os        string
	Platform  string
	Version   Version
	Extension string
}

type Configuration struct {
	Version    Version
	Timeout    time.Duration
	Net        string
	CmdOptions []string
	PidFile    string
	Username   string
	Password   string
	DBName     string
}

func Extract(d *Distribution, command Command) {
	workDir :=fmt.Sprintf("%v/%v/", d.Dir, d.Os)
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

	for _, file := range r.File {
		name := regexFilename.ReplaceAllString(file.Name, "$1")
		isExec := regexCommand.MatchString(name)
		fmt.Printf("%v\t\t\t%v\n", file.Name, name)
		if isExec {
			extractOneFile(file, workDir, name)
		}
	}
}

func extractOneFile(file *zip.File, workDir string, name string) {
	reader, _ := file.Open()
	//TODO: catch error
	defer reader.Close()

	f, _ := os.OpenFile(workDir+name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeTemporary)
	//TODO: catch error
	io.Copy(f, reader)
}
