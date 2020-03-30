package go2fe

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
		StartDev()
		fmt.Println("前端进程已启动")
		return
	}
	if strings.Contains(args, buildFlag) {
		// 构建前端代码
		fmt.Println("构建前端代码")
		StartBuild()
		fmt.Println("前端代码构建完毕")
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
	initComd := exec.Command("npm", "i", "-S", "go2fe-node-template", "--registry=\"http://registry.npmjs.org/\"")
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}
	// 必须指定工作路径, 否则找不到对应文件
	initComd.Dir = clientPathURI
	initComd.Stdout = &stdout
	initComd.Stderr = &stderr
	// @todo(yaozeyuan) 暂时屏蔽, 方便debug后续代码
	fmt.Println("command => ", initComd.String())
	fmt.Println("安装模板代码包")
	initComd.Run()
	fmt.Println("模板代码包安装完毕")

	// 释放前端模板
	uppackageFeComd := exec.Command("node", "bootstrap.js")
	uppackageFeComd.Dir = clientPathURI
	fmt.Println("准备释放前端模板代码")
	uppackageFeComd.Run()
	fmt.Println("前端代码释放完毕")

	// 进入client目录, 执行npm i
	fmt.Println("执行npm install")
	npmComd := exec.Command("npm", "i", "--registry=\"http://registry.npmjs.org/\"")
	npmComd.Dir = clientPathURI
	npmComd.Run()
	fmt.Println("npm install执行完毕")
	return
}

// StartBuild 执行构建(dev命令使用脚本启动会更好)
func StartBuild() {
	currentPath := getCurrentPath()
	clientPathURI := filepath.Join(currentPath, "client")
	isExist := isPathExist(clientPathURI)
	if isExist == false {
		fmt.Println("未检测到client文件夹,请先执行go run main.go go2fe:init 初始化项目")
		return
	}

	npmComd := exec.Command("npm", "run", "build")
	npmComd.Dir = clientPathURI
	fmt.Println("执行npm run build, 构建前端代码")
	npmComd.Run()
	fmt.Println("前端代码构建完毕")
	return
}

// StartDev 执行构建(dev命令使用脚本启动会更好)
func StartDev() {
	currentPath := getCurrentPath()
	npmComd := exec.Command("npm", "run", "dev")
	clientPathURI := filepath.Join(currentPath, "client")
	isExist := isPathExist(clientPathURI)
	if isExist == false {
		fmt.Println("未检测到client文件夹,请先执行go run main.go go2fe:init 初始化项目")
		return
	}
	npmComd.Dir = clientPathURI
	fmt.Println("执行npm run dev, 启动开发环境")
	npmComd.Run()
	return
}
