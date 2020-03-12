package generate

import (
	"encoding/json"
	"fmt"
	config "go2fe/democonfig"

	// 通过反射拿到数据模型结构
	"reflect"
)

// type TypeParseResult struct {
// 	Project     string          `json:"project"`
// 	CountDateAt int             `json:"countDateAt"`
// 	Detail      []TypeUILibItem `json:"detail"`
// }

// 数据模型-key模型
type dataKeyModel struct {
	Key          string `json:"key"`
	KeyType      string `json:"var_type"`
	IsShowInList bool   `json:"is_show_in_list"`
	IsUniqueKey  bool   `json:"is_unique_key"`
}

type dataModel struct {
	// 定义每一字段的key
	KeyList []dataKeyModel `json:"key_list"`
}

type typeJSONConfig struct {
	DataModel dataModel `json:"data_model"`
	// 一个数据模型可以对应多个page页面, true则自动生成, false则忽略
	// list页中, 根据page定义, 带有增加/删除/修改/详情按钮
	// list一定为true
	Page config.Page `json:"page_config"`
	// 页面基础地址(前端自动归并, 并补全出list/create/update路径)
	BasePath string `json:"base_url_path"`
	// 接口基础地址
	BaseAPIPath string `json:"base_api_path"`
}

// GetJSONConfig 输出项目json配置
func GetJSONConfig() (jsonConfigListJSONStr string) {
	pageList := config.PageList
	var jsonConfigList []typeJSONConfig
	for _, page := range pageList {
		var jsonConfig typeJSONConfig
		jsonConfig.BaseAPIPath = page.BaseAPIPath
		jsonConfig.BasePath = page.BasePath
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
				continue
			}

			keyType := field.Type.String()
			keyModel.KeyType = keyType
			_, keyModel.IsUniqueKey = field.Tag.Lookup("unique_key")
			_, keyModel.IsShowInList = field.Tag.Lookup("show")

			// 将数据结构打到jsonConfig中
			jsonConfig.DataModel.KeyList = append(jsonConfig.DataModel.KeyList, keyModel)
			fmt.Println("第", fieldIndex+1, "个属性:", keyModel)
			fieldIndex = fieldIndex + 1
		}
		jsonConfigList = append(jsonConfigList, jsonConfig)

	}

	jsonConfigListJSONBuf, _ := json.Marshal(jsonConfigList)
	jsonConfigListJSONStr = string(jsonConfigListJSONBuf)

	fmt.Println("pageList =>", pageList)
	fmt.Println("----------------")
	fmt.Println("json str =>", jsonConfigListJSONStr)
	return jsonConfigListJSONStr
}
