// +build windows,386
package env

func init() {
	MONGO_BITSIZE = "i386"
	MONGO_OS = "win32"
	MONGO_URL = "https://downloads.mongodb.org/"
	MONGO_EXT = "zip"
}
