// +build windows,amd64
package env

func init() {
	MONGO_BITSIZE = "x86_64"
	MONGO_OS = "win32"
	MONGO_EXT = "zip"
}
