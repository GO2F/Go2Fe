package go2fe

import (
	"fmt"
	"testing"
)

// CompontentModel demo数据模型
type CompontentModel struct {
	ID          string `json:"id" unique_key:"" show:"" title:"id"`
	DisplayName string `json:"display_name" show:"" title:"组件名"`
	PackageName string `json:"package_name" show:"" title:"包名"`
	DevListJSON string `json:"dev_list_json" show:"" title:"开发者"`
	Description string `json:"description" show:"" title:"描述"`
	SiteURL     string `json:"site_url" show:"" title:"网站主页"`
	Remark      string `json:"remark" title:"备注"`
}

// Run 启动测试
func Run() {
	fmt.Println("启动测试")
	fmt.Println("注册数据模型")
	model := ModelDefine{
		DataModel: CompontentModel{},
		Page: Page{
			Create: true,
			Update: false,
			Detail: true,
		},
		Name:        "测试模型",
		BasePath:    "/compontent",
		BaseAPIPath: "/api/compontent",
	}
	RegModel(model)
	fmt.Println("获取页面配置列表")
	configStr := GetJSONConfig()
	fmt.Println("页面配置数据为:", configStr)
	fmt.Println("输出配置")
	WriteConfig()
}

// TestHello 第一个测试
func TestHello(t *testing.T) {
	fmt.Println("success!")

	t.Log("start run")
	Run()
	t.Log("success 2")

}
