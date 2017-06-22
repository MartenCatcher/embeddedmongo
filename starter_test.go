package embeddedmongo

import (
	"testing"
	"os"
	"log"
	"regexp"
)

func TestNewProcess(T *testing.T) {

	wd, _ := os.Getwd()

	d := NewDistribution(Configuration{
		Version: V3_4_1,
		Dir:     wd + "/test/resources/",
	})

	path, err := Download(d)
	if path == "" {
		log.Printf("Download error: %v\n", err)
		return
	}

	log.Printf("Extracting: %v\n", path)

	files, err := Extract(d)
	if err != nil {
		log.Printf("Extract error: %v\n", err)
		return
	}

	log.Println("Extracted succssfully.")

	binary := ""
	re := regexp.MustCompile("mongod(.exe)?$")
	for _, file := range files {
		if re.FindString(file) != "" {
			binary = file
			break
		}
	}

	if binary == "" {
		log.Printf("Error: can't find the mongod binary in the extracted files.")
		return
	}

	p, err := NewProcess(binary, GetTmpDir(d))
	if err != nil {
		log.Printf("Executing error: %v\n", err)
		return
	}
	defer p.Stop()

	log.Printf("%+v", p)
}
