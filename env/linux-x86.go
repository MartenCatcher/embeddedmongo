// +build linux,386
package env

func init() {
	MONGO_BITSIZE = "i686"
	MONGO_OS = "linux"
	MONGO_EXT = "tgz"
}