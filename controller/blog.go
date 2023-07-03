package controller

import (
	"FurballCommunity_backend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
			"msg":  "成功修改文章！",
			"data": blog,
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
			"msg":  "成功获取blog！",
			"data": blog,
		})
	}
}

func isInArray(element uint, array []uint) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}

// LikeBlog
// @Summary 点赞博客（暂不可用）
// @Description 通过用户id和博客id完成点赞博客操作
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param   userid    query    uint     true      "user_id"
// @Param   blogid    query    uint     true      "blog_id"
// @Success 200 {string} string	"ok"
// @Router /v1/blog/like [PUT]
func LikeBlog(c *gin.Context) {
	userId := c.Query("userid")
	conv_userId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的用户id"})
		return
	}
	blogID := c.Query("blogid")
	conv_blogID, err := strconv.ParseInt(blogID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的博客id"})
		return
	}
	likedBlog, err := models.GetUserLikedBlog(uint(conv_userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}
	log.Println(likedBlog)
	if isInArray(uint(conv_blogID), likedBlog) {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "不能重复点赞"})
		return
	} else {
		blog, err := models.GetBlogInfo(uint(conv_blogID))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
			return
		}
		e := c.BindJSON(&blog)
		if e != nil {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": e.Error()})
			return
		}
		if err := models.LikeBlog(blog); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		}
		user, err := models.GetUserById(uint(conv_userId))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
			return
		}
		ee := c.BindJSON(&user)
		if ee != nil {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": ee.Error()})
			return
		}
		if err := models.AddToUserLikedList(user, uint(conv_blogID)); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": reStatusSuccess,
				"msg":  "点赞成功！",
			})
		}

	}
}

// DeleteBlog
// @Summary 删除博客
// @Description 通过blog_id删除博客
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param   blog_id    path    uint     true      "blog_id"
// @Success 200 {string} string	"ok"
// @Router /v1/blog/delete/{blog_id} [delete]
func DeleteBlog(c *gin.Context) {
	id, ok := c.Params.Get("blog_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	blogId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	if err := models.DeleteBlog(uint(blogId)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"msg":  "删除成功",
		})
	}
}
