package page

// pageType 页面类型
type pageType int32

const (
	// List List 页面
	List pageType = 0
	// Update Update 页面
	Update pageType = 1
	// Create Create 页面
	Create pageType = 2
)
