package controller

import (
	"FurballCommunity_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateBlog
// @Summary 创建社区文章
// @Description 创建一个社区文章 eg：{"user_id":2, "title":"标题", "content": "hhh"}
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param   blog    body    string   true      "userid + title + content"
// @Success 200 {string} string	"ok"
// @Router /v1/blog/create [post]
func CreateBlog(c *gin.Context) {
	var blog models.Blog
	err := c.BindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}
	if err := models.CreateBlog(&blog); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "成功发布博文", "blog": blog})
	}
}

// GetBlogList
// @Summary 获取社区文章列表
// @Description 获取所有社区文章
// @Tags Blog
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /v1/blog/getBlogList [get]
func GetBlogList(c *gin.Context) {
	blogList, err := models.GetBlogList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":      1,
			"blog_list": blogList,
		})
	}
}

// GetBlogListOfUser
// @Summary 获取用户文章列表
// @Description 通过用户id获取用户文章列表
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "user_id"
// @Success 200 {string} string	"ok"
// @Router /v1/blog/getUserBlog/{id} [get]
func GetBlogListOfUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id！"})
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	blogList, e := models.GetBlogListOfUser(uint(userId))
	if e != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": e.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "查询成功", "user_blog": blogList})
}

// UpdateBlog
// @Summary 更新社区文章
// @Description 通过id，更新blog，包括内容和标题等 eg：{"title":"wangwang", "content":"hhh"}
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "id"
// @Param   blog    body    string     true      "new_blog_info"
// @Success 200 {string} string	"ok"
// @Router /v1/blog/info/{id} [put]
func UpdateBlog(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	blogId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	blog, err := models.GetBlogInfo(uint(blogId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}
	e := c.BindJSON(&blog)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": e.Error()})
		return
	}
	if err := models.UpdateBlog(blog); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"meg":  "成功修改文章！",
			"blog": blog,
		})
	}
}

// GetBlogInfo
// @Summary 获取文章信息
// @Description 通过blog_id获取文章信息
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "blog_id"
// @Success 200 {string} string	"ok"
// @Router /v1/blog/info/{id} [get]
func GetBlogInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	blogId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	blog, err := models.GetBlogInfo(uint(blogId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"meg":  "成功获取blog！",
			"blog": blog,
		})
	}
}
