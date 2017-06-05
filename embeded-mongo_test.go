package embeded_mongo

import (
	"testing"
)

func TestExtract(T *testing.T) {
	d := &Distribution{Configuration: Configuration{Dir: "./test/resources", Version: V3_4_1}, Os: "win32", Platform: "x86_64", Extension: "zip"}
	Extract(d, Mongod)
}

func TestFork(T *testing.T) {
	var app = "./test/resources/win32/mongod.exe"
	p, _ := NewProcess(app,  "--logpath", "E:\\tools\\mongo\\logs\\mongo.log", "--dbpath", "E:\\tools\\mongo\\db")

	defer p.Stop()
}