package embeddedmongo

import (
	"log"
	"os"
	"testing"
)

func TestExtract(T *testing.T) {

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

	for _, file := range files {
		log.Print(file)
	}
	log.Printf("Extracted: %v elements.\n", len(files))

	err = os.RemoveAll(GetTmpDir(d))
	if err != nil {
		log.Printf("Remove directory error: %v\n", err)
		return
	}

	log.Print("Temporary directory with extracted files are removed.")
}
