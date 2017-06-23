package embeddedmongo

import (
	"log"
	"os"
	"regexp"
	"testing"
)

func TestNewDistribution(T *testing.T) {
	log.Printf("%+v", NewDistribution(Configuration{}))
}

func TestIntegration(T *testing.T) {

	wd, _ := os.Getwd()

	d := NewDistribution(
		Configuration{
			Version: V3_4_1,
			Dir:     wd + "/test/resources/",
		}, "https://fastdl.mongodb.org/",
	)

	_, err := Download(d)
	if err != nil {
		log.Printf("Download error: %v\n", err)
		panic(err)
	}

	extracted, err := Extract(d)
	if err != nil {
		log.Printf("Extract error: %v\n", err)
		panic(err)
	}

	binary := ""
	re := regexp.MustCompile("mongod(.exe)?$")
	for _, file := range extracted {
		if re.FindString(file) != "" {
			binary = file
			break
		}
	}

	p, err := NewProcess(binary, GetTmpDir(d))
	if err != nil {
		log.Printf("Executing error: %v\n", err)
		panic(err)
	}

	log.Printf("%+v", p)

	defer p.Stop()
}
