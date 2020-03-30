package go2fe

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 命令行flag列表
const initFlag = "go2fe:init"
const devFlag = "go2fe:dev"
const buildFlag = "go2fe:build"

// Bootstrap 启动项目
// 根据系统参数不同, 执行不同命令
func Bootstrap() {
	args := strings.Join(os.Args[1:], "")
	if strings.Contains(args, initFlag) {
		fmt.Println("go2fe bootstrap...")
		// 初始化
		fmt.Println("初始化项目代码")
		InitFeTemplate()
		fmt.Println("项目代码初始化完毕")
		return
	}
	if strings.Contains(args, devFlag) {
		fmt.Println("go2fe bootstrap...")
		// 启动dev环境
		fmt.Println("进程以命令模式启动")
		StartDev()
		return
	}
	if strings.Contains(args, buildFlag) {
		fmt.Println("go2fe bootstrap...")
		// 构建前端代码
		fmt.Println("构建前端代码")
		StartBuild()
		return
	}
	return
}

// InitFeTemplate 首次生成前端项目代码
func InitFeTemplate() {
	fmt.Println("初始化前端部分代码")
	fmt.Println("执行时间大约需要10分钟, 请耐心等待")
	currentPath := getCurrentPath()
	// 创建client文件夹
	clientPathURI := filepath.Join(currentPath, "client")
	resetDir(clientPathURI)
	// 创建static文件夹
	staticPathURI := filepath.Join(currentPath, "static")
	resetDir(staticPathURI)

	bootstrapJsURI := filepath.Join(clientPathURI, "bootstrap.js")
	// 写入bootstrap.js文件
	ioutil.WriteFile(bootstrapJsURI, []byte("var unpackage=require('go2fe-node-template');\nunpackage.default();"), 0777)

	// 安装包依赖
	fmt.Println("安装模板代码包")
	runCommand([]string{"npm", "i", "-S", "go2fe-node-template", "--registry=\"http://registry.npmjs.org/\""}, clientPathURI)
	fmt.Println("模板代码包安装完毕")

	// 释放前端模板
	fmt.Println("准备释放前端模板代码")
	runCommand([]string{"node", "bootstrap.js"}, clientPathURI)
	fmt.Println("前端代码释放完毕")

	// 进入client目录, 执行npm i
	fmt.Println("执行npm install")
	runCommand([]string{"npm", "i", "--registry=\"http://registry.npmjs.org/\""}, clientPathURI)
	fmt.Println("npm install执行完毕")
	return
}

func startCommand(env string) {
	currentPath := getCurrentPath()
	clientPathURI := filepath.Join(currentPath, "client")
	isExist := isPathExist(clientPathURI)
	if isExist == false {
		fmt.Println("未检测到client文件夹,请先执行go run main.go go2fe:init 初始化项目")
		return
	}

	// 检查页面配置
	fmt.Println("检查页面配置")
	modelList := GetModelList()
	fmt.Println("当前共配置页面", len(modelList), "个")
	if len(modelList) == 0 {
		fmt.Println("当前没有可用页面配置, 请先注册页面后再用")
		return
	}
	// 输出页面配置列表
	fmt.Println("更新页面配置")
	WriteConfig()
	fmt.Println("页面配置更新完毕")

	if env == "npm-run-dev" {
		fmt.Println("执行npm run dev, 启动开发环境")
		runCommand([]string{"npm", "run", "dev"}, clientPathURI)
	} else {
		fmt.Println("重置static文件夹")
		// 创建static文件夹
		staticPathURI := filepath.Join(currentPath, "static")
		resetDir(staticPathURI)
		fmt.Println("static文件夹重置完毕")
		fmt.Println("执行npm run build, 构建前端代码")
		runCommand([]string{"npm", "run", "build"}, clientPathURI)
		fmt.Println("前端代码构建完毕")
	}
	return
}

// StartBuild 执行构建(dev命令使用脚本启动会更好)
func StartBuild() {
	startCommand("npm-run-build")
	return
}

// StartDev 执行构建(dev命令使用脚本启动会更好)
func StartDev() {
	startCommand("npm-run-dev")
	return
}
