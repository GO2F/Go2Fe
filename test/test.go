package test

import (
	"fmt"

	// "github.com/GO2F/Go2Fe/generate"
	// "github.com/GO2F/Go2Fe/register"
	"github.com/GO2F/Go2Fe/go2fe"
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
	model := go2fe.DefineModel{
		DataModel: CompontentModel{},
		Page: go2fe.Page{
			Create: true,
			Update: false,
			Detail: true,
		},
		Name: "测试模型"
		BasePath:    "/compontent",
		BaseAPIPath: "/api/compontent",
	}
	go2fe.RegModel(model)
	fmt.Println("获取页面配置列表")
	configStr := go2fe.GetJSONConfig()
	fmt.Println("页面配置数据为:", configStr)
	fmt.Println("输出配置")
	go2fe.WriteConfig()
}
