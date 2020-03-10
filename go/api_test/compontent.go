package compontent

import (
	"net/http"
	"strconv"

	config "go2fe/demo-config"

	"github.com/gin-gonic/gin"
)

// Get 获取单条记录
func Get(c *gin.Context) {
	type RequestParam struct {
		// 正常使用id
		ID uint `form:"id"`
	}
	var params RequestParams
	c.BindQuery(&params)
	compontent := mCompontent.Get(params.ID)
	if compontent.ID == 0 {
		c.JSON(http.StatusOK, base.Failed("项目不存在", 2))
		return
	}
	c.JSON(http.StatusOK, base.Success(compontent))
	return
}

// GetList 获取列表
func GetList(c *gin.Context) {
	type RequestParams struct {
		Page uint `form:"page"`
	}
	page := c.DefaultQuery("page", "0")
	pageNum, _ := strconv.Atoi(page)
	compontentList := mCompontent.GetList(pageNum)
	c.JSON(http.StatusOK, base.Success(compontentList))
	return
}

// Apply 项目申请
func Apply(c *gin.Context) {
	// create中使用在demo-config中定义的数据结构
	type RequestParams struct {
		DisplayName string `json:"display_name"`
		PackageName string `json:"package_name"`
		DevListJSON string `json:"dev_list_json"`
		Description string `json:"description"`
		SiteURL     string `json:"site_url"`
		Remark      string `json:"remark"`
	}
	var params RequestParams
	c.BindJSON(&params)
	rawUser, isExist := c.Get("user")
	if isExist != true {
		c.JSON(http.StatusOK, base.Failed("请先登录", 2))
		return
	}
	user := rawUser.(mUser.User)

	compontent := mCompontent.Compontent{
		DisplayName: params.DisplayName,
		PackageName: params.PackageName,
		DevListJSON: params.DevListJSON,
		Description: params.Description,
		ApplyUcid:   user.Ucid,
		SiteURL:     params.SiteURL,
		Remark:      params.Remark,
	}

	existProj := mCompontent.GetUnscopedByPackageName(params.PackageName)
	if existProj.ID != 0 {
		if existProj.DeletedAt == nil {
			// 已存在
			c.JSON(http.StatusOK, base.Success(existProj))
			return
		} else {
			c.JSON(http.StatusOK, base.Failed("项目 "+existProj.PackageName+" 已被删除,请联系管理员恢复", 1))
			return
		}

	}

	mCompontent.Create(compontent)
	// 发送邮件通知
	uMail.SendMail(config.Mail.NotifyList, "排行榜项目收到新组件库注册申请", "排行榜项目收到新组件库注册申请, packageName =>"+params.PackageName+", 请及时处理")

	c.JSON(http.StatusOK, base.Success(params))
	return
}

// Approve 项目审核
func Approve(c *gin.Context) {
	rawUser, isExist := c.Get("user")
	if isExist == false {
		c.JSON(http.StatusOK, base.Login())
		return
	}
	user := rawUser.(mUser.User)
	if user.IsAdmin == 0 {
		c.JSON(http.StatusOK, base.Failed("只有管理员有权限批准项目", 1))
		return
	}
	type RequestParams struct {
		ID uint `json:"ID"`
	}
	var params RequestParams
	c.BindJSON(&params)
	compontent := mCompontent.Get(params.ID)
	if compontent.ID == 0 {
		c.JSON(http.StatusOK, base.Failed("项目不存在", 2))
		return
	}
	compontent.IsAllow = 1
	mCompontent.Update(compontent)

	c.JSON(http.StatusOK, base.Success(compontent))
	return
}

// Update 更新项目数据
func Update(c *gin.Context) {
	type RequestParams struct {
		ID          uint   `json:"ID"`
		DisplayName string `json:"displayName"`
		DevListJSON string `json:"devListJson"`
		Description string `json:"description"`
		SiteURL     string `json:"site"`
		Remark      string `json:"remark"`
	}
	var params RequestParams
	c.BindJSON(&params)
	compontent := mCompontent.Get(params.ID)
	if compontent.ID == 0 {
		c.JSON(http.StatusOK, base.Failed("项目不存在", 2))
		return
	}
	if params.DisplayName != "" {
		compontent.DisplayName = params.DisplayName
	}
	if params.DevListJSON != "" {
		compontent.DevListJSON = params.DevListJSON
	}
	if params.Description != "" {
		compontent.Description = params.Description
	}
	if params.SiteURL != "" {
		compontent.SiteURL = params.SiteURL
	}
	if params.Remark != "" {
		compontent.Remark = params.Remark
	}

	mCompontent.Update(compontent)

	c.JSON(http.StatusOK, base.Success(compontent))
	return
}

// Remove 删除项目
func Remove(c *gin.Context) {
	rawUser, isExist := c.Get("user")
	if isExist == false {
		c.JSON(http.StatusOK, base.Login())
		return
	}
	type RequestParams struct {
		ID uint `form:"id"`
	}
	var params RequestParams
	c.BindQuery(&params)
	compontent := mCompontent.Get(params.ID)
	if compontent.ID == 0 {
		c.JSON(http.StatusOK, base.Failed("项目不存在", 2))
		return
	}
	user := rawUser.(mUser.User)
	if user.IsAdmin == 0 || compontent.ApplyUcid != user.Ucid {
		c.JSON(http.StatusOK, base.Failed("只有创建者/管理员有权限删除项目", 1))
		return
	}
	mCompontent.Delete(compontent.ID)

	c.JSON(http.StatusOK, base.Success("删除成功"))
	return
}
