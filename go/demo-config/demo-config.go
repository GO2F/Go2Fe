package demo-config
import ("go2fe/pattern/page")

// 配置demo, 未来使用该文件作为配置模板, 生成接口规范&前端定义

const 

struct QueryParam {
	name string "参数名",
	paramType string "参数类型, 只能为go的基本类型",
}
struct RequestParam {
	name string "参数名",
	paramType string "参数类型, 只能为go的基本类型",
}

// DataModel 数据模型, 用于定义操作的数据结构
struct DataModel {
	name string "参数名",
	paramType string "参数类型, 只能为go的基本类型",
}


struct ListPage {
	pageType page `页面类型, 枚举值, List/Update/Create 三者之一`,
	queryPath string "/compontent/list",
	queryParam QueryParam[] `query参数列表`,
	requestType string `第一期只支持json类型`
	requestParam RequestParam[] `request参数列表`,
	dataModel  DataModel `数据模型, 用于生成list列表/form表单`
}


