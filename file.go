package utils

import (
	"os"
)

//文件、路径是否存在
func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
