package embeddedmongo

import (
	"log"
	"os"
	"testing"
)

func TestExtract(T *testing.T) {
	d := &Distribution{Configuration: Configuration{Dir: "./test/resources", Version: V3_4_1}, Os: "win32", Platform: "x86_64", Extension: "zip"}
	Extract(d, Mongod)
}

func TestFork(T *testing.T) {
	p, _ := NewProcess("mongod.exe", "./test/resources/win32/")

	defer p.Stop()
}

func TestNewDistribution(T *testing.T) {
	log.Printf("%+v", NewDistribution(Configuration{}))
}

func TestIntegration(T *testing.T) {
	wd, _ := os.Getwd()
	d := NewDistribution(Configuration{Version: V3_4_1, Dir: wd + "/test/resources/"})
	err := Download(GetDistributionName(d), GetWorkDir(d), GetDistributionUrl(d))
	if err != nil {
		log.Printf("Download error: %v\n", err)
		panic(err)
	}

	app, err := Extract(d, Mongod)
	if err != nil {
		log.Printf("Extracting error: %v\n", err)
		panic(err)
	}

	p, err := NewProcess(app, GetTmpDir(d))
	if err != nil {
		log.Printf("Extracting error: %v\n", err)
		panic(err)
	}

	log.Printf("%+v", p)

	defer p.Stop()
}
