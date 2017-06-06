// +build darwin,386
package env

func init() {
	MONGO_BITSIZE = "i386"
	MONGO_OS      = "osx"
	MONGO_URL     = "https://fastdl.mongodb.org/"
	MONGO_EXT     = "tgz"
}