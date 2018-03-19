package utils

import "os"

func PathOrFileExist(_path string, isDir bool) bool {
	stat, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return (!stat.IsDir() && !isDir) || (stat.IsDir() && isDir)
}