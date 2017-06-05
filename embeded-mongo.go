package embeded_mongo

import (
	"runtime"
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
	Configuration
	Url       string
	Os        string
	Platform  string
	Extension string
}

type Configuration struct {
	Version Version
	Dir     string
}

func NewDistribution(configuration Configuration) *Distribution {
	//TODO: use build instructions
	var extension string
	var url string
	if runtime.GOOS == "windows" {
		extension = "zip"
		url = "https://downloads.mongodb.org/"
	} else {
		extension = "tgz";
		url = "https://fastdl.mongodb.org/"
	}

	var os string
	switch runtime.GOOS {
	case "darwin":  os = "osx"
	case "freebsd": os = "freebsd"
	case "windows": os = "win32"
	case "linux":   os = "linux"
	}

	var bitSize string
	switch runtime.GOARCH {
	case "386":
		if runtime.GOOS == "linux" {
			bitSize = "i686"
		} else {
			bitSize = "i386"
		}
	case "amd64":
		bitSize = "x86_64"
	}

	// return distribution.getPlatform() == Platform.Windows?:;

	return &Distribution{Configuration: configuration, Url: url, Os: os, Platform: bitSize, Extension: extension}
}