package embeddedmongo

import (
	"log"
	"os"
	"testing"
)

func TestDownload(T *testing.T) {

	wd, _ := os.Getwd()

	d := NewDistribution(Configuration{
		Version: V3_4_1,
		Dir:     wd + "/test/resources/",
	})

	file, err := Download(d)
	if err != nil {
		if file == "" {
			log.Printf("Download error: %v\n", err)
		} else {
			log.Printf("Warning: %v (%v)\n", err, file)
		}
	} else {
		log.Printf("Downloaded: %v\n", file)
	}

}
