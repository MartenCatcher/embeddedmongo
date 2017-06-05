package embeded_mongo

import (
	"testing"
)

func TestExtract(T *testing.T) {
	d := &Distribution{Dir:"test/resources", Os:"win32", Platform:"x86_64", Version:V3_4_1, Extension:"zip"}
	Extract(d, Mongod)
}
