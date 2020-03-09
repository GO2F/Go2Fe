package main

import (
	"Go2Fe/config"
	"fmt"
)

func main() {
	fmt.Println(config.GetSchoolMenber())

	/*
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		goS := c.Query("Go")
		feS := c.Query("Fe")
		c.String(http.StatusOK, "Hello %s %s", goS, feS)
	})

	r.POST("/post", func(c *gin.Context) {
		name := c.PostForm("name")                            //找不到name直接返回0值
		password := c.DefaultPostForm("password", "00000000") //找不到password赋默认值
		result, ok := c.GetPostForm("ok")                     //判断是否能找到，找不到返回false
		fmt.Println(ok)
		c.String(http.StatusOK, "hello %s %s %s", name, password, result)
	})

	r.Run(":9999") // listen and serve on 0.0.0.0:8080
	*/
}
