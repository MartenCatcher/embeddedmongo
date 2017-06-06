package embeded_mongo

import (
	"embeded-mongo/env"
	"fmt"
	"os"
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
	return &Distribution{Configuration: configuration, Url: env.MONGO_URL, Os: env.MONGO_OS, Platform: env.MONGO_BITSIZE, Extension: env.MONGO_EXT}
}

func GetDistributionName(d *Distribution) string {
	return fmt.Sprintf("%v-%v-%v-%v.%v", "mongodb", d.Os, d.Platform, d.Version, d.Extension)
}

func GetDistributionUrl(d *Distribution) string {
	return fmt.Sprintf("%v%s/%v", d.Url, d.Os, GetDistributionName(d))
}

func GetWorkDir(d *Distribution) string {
	return fmt.Sprintf("%v%v/", d.Dir, d.Os)
}

func CreateWorkDir(path string) (error) {
	err := os.MkdirAll(path, 0755)
	return err
}
