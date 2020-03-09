package pattern

// PageType 页面类型
type PageType int32

const (
	// PageList List 页面
	List PageType = 0
	// PageUpdate Update 页面
	Update PageType = 1
	// PageCreate Create 页面
	Create PageType = 2
)
