// +build linux,386
package env

func init() {
	MONGO_BITSIZE = "i686"
	MONGO_OS = "linux"
	MONGO_URL = "https://fastdl.mongodb.org/"
	MONGO_EXT = "tgz"
}