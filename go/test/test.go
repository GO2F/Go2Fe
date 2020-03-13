package test

import (
	"fmt"
	"go2fe/generate"
	"go2fe/register"
)

// CompontentModel demo数据模型
type CompontentModel struct {
	ID          string `json:"id" unique_key:""`
	DisplayName string `json:"display_name" show:""`
	PackageName string `json:"package_name"`
	DevListJSON string `json:"dev_list_json"`
	Description string `json:"description"`
	SiteURL     string `json:"site_url"`
	Remark      string `json:"remark"`
}

// Run 启动测试
func Run() {
	fmt.Println("启动测试")
	fmt.Println("注册数据模型")
	model := register.DefineModel{
		DataModel: CompontentModel{},
		Page: register.Page{
			Create: true,
			Update: false,
			Detail: true,
		},
		BasePath:    "/compontent",
		BaseAPIPath: "/api/compontent",
	}
	register.Register(model)
	fmt.Println("获取页面配置列表")
	configStr := generate.GetJSONConfig()
	fmt.Println("页面配置数据为:", configStr)
	fmt.Println("输出配置")
	generate.WriteConfig()
}
