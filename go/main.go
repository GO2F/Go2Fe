package main

import (
	// "fmt"
	// "reflect"
	"go2fe/generate"
)

// RequestParams test
type RequestParams struct {
	DisplayName string `json:"display_name"`
	PackageName string `json:"package_name"`
	DevListJSON string `json:"dev_list_json"`
	Description string `json:"description"`
	SiteURL     string `json:"site_url"`
	Remark      string `json:"remark"`
}

func main() {

	generate.Test()

	// a := RequestParams{
	// 	DisplayName: "123123",
	// 	PackageName: "123",
	// }

	// // 此处必须使用引用
	// test := reflect.TypeOf(&a)
	// test1 := test.Elem()
	// test2, _ := test1.FieldByName("DevListJSON")
	// test3 := test2.Tag
	// goType := test2.Type              // 可以拿到数据类型string
	// test4 := test3.Get("json")        // 可以拿到json标记对应的值: display_name
	// notExist := test3.Get("notExist") // 可以拿到json标记对应的值: display_name
	// fmt.Println("tag => ", test)
	// fmt.Println("tag => ", test1)
	// fmt.Println("tag => ", test2)
	// fmt.Println("tag => ", test3)
	// fmt.Println("goType => ", goType)
	// fmt.Println("tag => ", test4)
	// fmt.Println("notExist tag => ", notExist)
	return
	// 	test1 := test.Elem()
	// 	test2, _ := test1.FieldByName("DisplayName")
	// 	test3 := test2.Tag
	// 	tag := test3.Get("json")
	// 	fmt.Println("tag => ", tag)
	// 	fmt.Println("hello world!\n")
	// 	fmt.Println(os.Args)
}
