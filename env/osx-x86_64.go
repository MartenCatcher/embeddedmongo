// +build darwin,386
package env

func init() {
	MONGO_BITSIZE = "x86_64"
	MONGO_OS      = "osx"
	MONGO_EXT     = "tgz"
}