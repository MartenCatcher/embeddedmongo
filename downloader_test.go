package embeded_mongo

import (
	"testing"
	"log"
)

func TestDownload(T *testing.T) {
	d := NewDistribution(Configuration{Version:V3_4_1, Dir:"./test/resources/"})
	err := Download(GetDistributionName(d), GetWorkDir(d), GetDistributionUrl(d))
	if err != nil {
		log.Printf("Error ocurs: %v\n", err)
	}
}