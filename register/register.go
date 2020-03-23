package register

// Page 页面配置
type Page struct {
	Create bool `json:"create"`
	Update bool `json:"update"`
	Detail bool `json:"detail"`
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

// Register 由业务方使用, 注册数据模型
func Register(model DefineModel) {
	pageList = append(pageList, model)
	return
}

// GetModelList 由github.com/GO2F/Go2Fe使用, 用于获取所有数据模型列表, 生成配置
func GetModelList() []DefineModel {
	return pageList
}
