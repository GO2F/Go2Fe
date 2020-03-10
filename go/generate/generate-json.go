package generate

import (
	"fmt"
	config "go2fe/democonfig"
)

// Test 测试
func Test() {
	pageList := config.PageList
	fmt.Println("pageList =>", pageList)
}
