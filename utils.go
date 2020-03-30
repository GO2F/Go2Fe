package go2fe

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

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

func isPathExist(pathURI string) (isExist bool) {
	_, err := os.Stat(pathURI)
	if err == nil {
		return true
	}
	return false
}

func runCommand(argv []string, dir string) bool {
	cmd := exec.Command(argv[0], argv[1:]...)
	cmd.Dir = dir
	fmt.Println("执行命令:", cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}
	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}
