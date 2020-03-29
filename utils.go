package go2fe

import "os"

func getCurrentPath() string {
	currentPath, _ := os.Getwd()
	return currentPath
}

// resetDir 删除文件夹下所有文件, 并重新创建文件夹
func resetDir(targetPath string) (isSuccess bool) {
	if len(targetPath) < 3 {
		// 路径长度不对
		return false
	}
	os.RemoveAll(targetPath)
	os.MkdirAll(targetPath, os.ModePerm)
	return true
}
