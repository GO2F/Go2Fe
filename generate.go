package go2fe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	// 通过反射拿到数据模型结构
	"reflect"

	"go/build"
	"os"
	"os/exec"
)

const version = 0.1

// 发布后的包名
const packageName = "github.com/GO2F/Go2Fe"

// type TypeParseResult struct {
// 	Project     string          `json:"project"`
// 	CountDateAt int             `json:"countDateAt"`
// 	Detail      []TypeUILibItem `json:"detail"`
// }

// 数据模型-key模型
type dataKeyModel struct {
	Key string `json:"key"`
	// 字段中文名
	Title        string `json:"title"`
	KeyType      string `json:"var_type"`
	IsShowInList bool   `json:"is_show_in_list"`
	IsUniqueKey  bool   `json:"is_unique_key"`
}

type dataModel struct {
	// 定义每一字段的key
	KeyList []dataKeyModel `json:"key_list"`
}

type typeJSONConfig struct {
	Version   float32   `json:"version" desc:"配置版本号,使用一位小数作为配置版本"`
	DataModel dataModel `json:"data_model"`
	// 一个数据模型可以对应多个page页面, true则自动生成, false则忽略
	// list页中, 根据page定义, 带有增加/删除/修改/详情按钮
	// list一定为true
	Page Page `json:"page_config"`
	// 数据模型在前端展示的名称
	Name string `json:"name"`
	// 页面基础地址(前端自动归并, 并补全出list/create/update路径)
	BasePath string `json:"base_url_path"`
	// 接口基础地址
	BaseAPIPath string `json:"base_api_path"`
}

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

// InitClientTemplate 生成移动端代码
func InitClientTemplate() {
	currentPath := getCurrentPath()
	// 创建client文件夹
	clientPathURI := currentPath + "/client"
	resetDir(clientPathURI)
	// 创建static文件夹
	staticPathURI := currentPath + "/static"
	resetDir(staticPathURI)

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	// 将resource/node_template下的文件复制到client中
	selfPath := gopath + "/pkg/mod" + "/" + packageName
	selfTemplatePath := selfPath + "/resource/node_template"
	comd := exec.Command("cp", "-r", selfTemplatePath, clientPathURI)
	// 必须指定工作路径, 否则找不到对应文件
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}
	// comd.Dir = nodeRuntimePath
	comd.Stdout = &stdout
	comd.Stderr = &stderr
	// @todo(yaozeyuan) 暂时屏蔽, 方便debug后续代码
	fmt.Println("command => ", comd.String())
	fmt.Println("将前端模板释放到client文件夹中")
	comd.Run()
	fmt.Println("释放完毕")
	fmt.Println("Stdout:", string(stdout.Bytes()))
	fmt.Println("Stderr:", string(stderr.Bytes()))
	// 进入client目录, 执行npm i
	fmt.Println("执行npm install")
	npmComd := exec.Command("npm", "i")
	npmComd.Dir = clientPathURI
	npmComd.Run()
	fmt.Println("npm install执行完毕")
	return
}

// GetJSONConfig 输出项目json配置
func GetJSONConfig() (jsonConfigListJSONStr string) {
	modelList := GetModelList()
	var jsonConfigList []typeJSONConfig
	for _, page := range modelList {
		var jsonConfig typeJSONConfig
		jsonConfig.Version = version
		jsonConfig.BaseAPIPath = page.BaseAPIPath
		jsonConfig.BasePath = page.BasePath
		jsonConfig.Name = page.Name
		jsonConfig.Page = page.Page

		dataModel := page.DataModel

		// 通过反射拿到所有数据结构
		// 这里可以直接使用dataModel, 而不是dataModel的地址, 使用dataModel地址拿到的是DefineModel中定义的数据类型(interface{}), 而不是
		dataModelType := reflect.TypeOf(dataModel)
		totalFieldNum := dataModelType.NumField()
		fieldIndex := 0
		for fieldIndex < totalFieldNum {
			var keyModel = dataKeyModel{}
			field := dataModelType.Field(fieldIndex)
			keyModel.Key = field.Tag.Get("json")
			if keyModel.Key == "" {
				// 如果没定义key值, 不应该输出, 而是直接跳过
				continue
			}

			keyType := field.Type.String()
			keyModel.KeyType = keyType
			_, keyModel.IsUniqueKey = field.Tag.Lookup("unique_key")
			_, keyModel.IsShowInList = field.Tag.Lookup("show")
			_, keyModel.Title = field.Tag.Get("title")

			// 将数据结构打到jsonConfig中
			jsonConfig.DataModel.KeyList = append(jsonConfig.DataModel.KeyList, keyModel)
			fmt.Println("第", fieldIndex+1, "个属性:", keyModel)
			fieldIndex = fieldIndex + 1
		}
		jsonConfigList = append(jsonConfigList, jsonConfig)

	}

	jsonConfigListJSONBuf, _ := json.Marshal(jsonConfigList)
	jsonConfigListJSONStr = string(jsonConfigListJSONBuf)

	fmt.Println("pageList =>", modelList)
	fmt.Println("----------------")
	fmt.Println("json str =>", jsonConfigListJSONStr)
	return jsonConfigListJSONStr
}

// WriteConfig 写入配置
func WriteConfig() {
	configStr := GetJSONConfig()
	currentPath := getCurrentPath()
	targetPathURI := currentPath + "/client/src/config/go2fe_generate_config.js"
	ioutil.WriteFile(targetPathURI, []byte("export default  "+configStr), 0777)
}

// StartBuild 执行构建(dev命令使用脚本启动会更好)
func StartBuild() {
	currentPath := getCurrentPath()
	npmComd := exec.Command("npm", "run", "build")
	clientPathURI := currentPath + "/client"
	npmComd.Dir = clientPathURI
	fmt.Println("执行npm run build, 构建前端代码")
	npmComd.Run()
	fmt.Println("前端代码构建完毕")
	return
}
