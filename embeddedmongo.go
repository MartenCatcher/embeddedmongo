package embeddedmongo

import (
	"fmt"
	"os"

	"github.com/MartenCatcher/embeddedmongo/env"
	"github.com/MartenCatcher/embeddedmongo/uuid"
)

type (
	Command int
	Version string
)

const (
	V3_4_1 Version = "3.4.1"
	DEFAULT_URL = "https://fastdl.mongodb.org/"
)

type Distribution struct {
	Configuration
	Url       string
	Os        string
	Platform  string
	Extension string
	Tmp       string
}

type Configuration struct {
	Version Version
	Dir     string
}

func NewDistribution(configuration Configuration, url ...string) *Distribution {
	return &Distribution{
		Configuration: configuration,
		Url:           getUrl(url),
		Os:            env.MONGO_OS,
		Platform:      env.MONGO_BITSIZE,
		Extension:     env.MONGO_EXT,
		Tmp:           uuid.Generate().String(),
	}
}

func getUrl(url []string) string {
	if len(url) == 0 {
		return DEFAULT_URL
	}
	return url[0]
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

func GetTmpDir(d *Distribution) string {
	return fmt.Sprintf("%v%v/%v/", d.Dir, d.Os, d.Tmp)
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, 0755)
	return err
}
