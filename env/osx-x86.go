// +build darwin,386
package env

func init() {
	MONGO_BITSIZE = "i386"
	MONGO_OS      = "osx"
	MONGO_EXT     = "tgz"
}