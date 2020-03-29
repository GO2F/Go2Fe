package go2fe

import (
	"fmt"
	"os"
	"strings"
)

// 命令行flag列表
const initFlag = "go2fe:init"
const devFlag = "go2fe:dev"
const buildFlag = "go2fe:build"

// Bootstrap 启动项目
// 根据系统参数不同, 执行不同命令
func Bootstrap() {
	fmt.Println("bootstrap")
	args := strings.Join(os.Args[1:], "")
	if strings.Contains(args, initFlag) {
		// 初始化
		fmt.Println("初始化项目代码")
		InitFeTemplate()
		fmt.Println("项目代码初始化完毕")
		return
	}
	if strings.Contains(args, devFlag) {
		// 启动dev环境
		fmt.Println("进程以命令模式启动")
		return
	}
	if strings.Contains(args, buildFlag) {
		// 构建前端代码
		fmt.Println("进程以命令模式启动")
		return
	}
	return
}
