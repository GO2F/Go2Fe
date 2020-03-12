package democonfig

// import ("go2fe/pattern/page")

// // 配置demo, 未来使用该文件作为配置模板, 生成接口规范&前端定义
// struct QueryParam {
// 	name string "参数名",
// 	paramType string "参数类型, 只能为go的基本类型",
// }
// struct RequestParam {
// 	name string "参数名",
// 	paramType string "参数类型, 只能为go的基本类型",
// }

// // DataModel 数据模型, 用于定义操作的数据结构
// struct DataModel {
// 	name string "参数名",
// 	paramType string "参数类型, 只能为go的基本类型",
// }

// struct ListPage {
// 	// 第一期项目中
// 	// List接口默认要求支持自带page(默认1)/size(默认10)两个参数
// 	// 使用强约束, 避免额外定义
// 	pageType page `页面类型, 枚举值, List/Update/Create 三者之一`,
// 	queryPath string "/compontent/list",
// 	// queryParam QueryParam[] `第一期不支持query参数列表`,
// 	// requestType string `第一期只支持json类型`
// 	// requestParam RequestParam[] `request参数列表`,
// 	dataModel  DataModel `数据模型, 用于生成list列表/form表单`
// }

// const page = ListPage{
// 	pageType: page.List,
// 	queryPath: "/api/compontent/list",
// 	dataModel RequestParams
// }

// // 理想写法
// list := interface{}{
// 	// 想要一个list页
// 	page: "list",
// 	// 页面基础地址(系统自动归并, 并补全出list/create/update路径)
// 	basePath: "#/compontent"
// 	// 接口地址为
// 	queryPath: "/api/compontent/list",
// 	// 需要处理的数据模型为
// 	// 模型中有字段标记, 可以取得以下信息
// 	// 1.	展示的字段=> show: true/false
// 	// 2. 	每个字段对应的数据类型
// 	// 3.	每个字段的值(直接使用字段名, 反正不需要填写)
// 	// 4.	是否为需要特殊展示的字段(例如日期型timestamp / 勾选型 checkbox / etc)
// 	dataModel DataModel{}// 此处不能传入具体值, 而应该是数据类型, 或者基于数据类型随意生成的具体值
// 	// 接口响应值不需要声明, 走默认模式
// }
// // 理想写法
// create := interface{}{
//  	// 想要一个create页
// 	page: "create",
// 	// 页面基础地址(系统自动归并, 并补全出list/create/update路径)
// 	basePath: "#/compontent"
// 	// 接口地址为
// 	queryPath: "/api/compontent/create",
// 	// 需要处理的数据模型为
// 	// 模型中有字段标记, 可以取得以下信息
// 	// 1.	展示的字段=> show: true/false
// 	// 2. 	每个字段对应的数据类型
// 	// 3.	每个字段的值(直接使用字段名, 反正不需要填写)
// 	// 4.	是否为需要特殊展示的字段(例如日期型timestamp / 勾选型 checkbox / etc)
// 	dataModel DataModel{}// 此处不能传入具体值, 而应该是数据类型, 或者基于数据类型随意生成的具体值
// 	// 接口响应值不需要声明, 走默认模式
// }

// Page 页面配置
type Page struct {
	Create bool `json:"create"`
	Update bool `json:"update"`
	Detail bool `json:"detail"`
}

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

// DefineModel 数据模型定义
type DefineModel struct {
	// 首先定义数据模型
	// 需要处理的数据模型为
	// 模型中有字段标记, 可以取得以下信息
	// 1.	展示的字段=> show: true/false
	// 2. 	每个字段对应的数据类型
	// 3.	每个字段的值(直接使用字段名, 反正不需要填写)
	// 4.	是否为需要特殊展示的字段(例如日期型timestamp / 勾选型 checkbox / etc)
	// 5.	第一期, 不支持复合类型字段.
	// 此处不能传入具体值, 而应该是数据类型, 或者基于数据类型随意生成的具体值
	DataModel interface{}
	// 接口响应值不需要声明, 走默认模式

	// 一个数据模型可以对应多个page页面, true则自动生成, false则忽略
	// list页中, 根据page定义, 带有增加/删除/修改/详情按钮
	// list一定为true
	Page Page
	// 页面基础地址(前端自动归并, 并补全出list/create/update路径)
	BasePath string
	// 接口基础地址
	BaseAPIPath string
}

// PageList 用于存放所有页面记录
var pageList []DefineModel

// RegisterModel 由业务方使用, 注册数据模型
func RegisterModel(model DefineModel) {
	pageList = append(pageList, model)
	return
}

// GetModelList 由go2fe使用, 用于获取所有数据模型列表, 生成配置
func GetModelList() []DefineModel {
	return pageList
}

func init() {

	page := DefineModel{
		// 首先定义数据模型
		// 需要处理的数据模型为
		// 模型中有字段标记, 可以取得以下信息
		// 1.	展示的字段=> show: true/false
		// 2. 	每个字段对应的数据类型
		// 3.	每个字段的值(直接使用字段名, 反正不需要填写)
		// 4.	是否为需要特殊展示的字段(例如日期型timestamp / 勾选型 checkbox / etc)
		// 5.	第一期, 不支持复合类型字段.
		// 此处不需要填写具体值
		DataModel: CompontentModel{
			ID: "123",
		},
		// 接口响应值不需要声明, 走默认模式

		// 一个数据模型可以对应多个page页面, true则自动生成, false则忽略
		// list页中, 根据page定义, 带有增加/删除/修改/详情按钮
		// list一定为true
		Page: Page{
			// list: true,
			Create: true,
			Update: true,
			Detail: false,
		},
		// 页面基础地址(前端自动归并, 并补全出list/create/update路径)
		BasePath: "/compontent",
		// 接口基础地址
		BaseAPIPath: "/api/compontent/",
	}

	RegisterModel(page)
}
