package embeddedmongo

import (
	"log"
	"testing"
)

func TestDownload(T *testing.T) {
	d := NewDistribution(Configuration{Version: V3_4_1, Dir: "./test/resources/"})
	err := Download(GetDistributionName(d), GetWorkDir(d), GetDistributionUrl(d))
	if err != nil {
		log.Printf("Download error: %v\n", err)
		panic(err)
	}
}
