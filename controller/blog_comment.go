package controller

import (
	"FurballCommunity_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateBlogComment
// @Summary 创建博客评论
// @Description 根据用户id、博客id，添加一个新的评论 eg：{ "user_id":1, "blog_id":2, "content":"文章很有用" }
// @Tags BlogCmt
// @Accept  json
// @Produce  json
// @Param   blogCmt    body    string   true      "userid + blogId + content"
// @Success 200 {string} string	"ok"
// @Router /v1/BlogCmt/create [post]
func CreateBlogComment(c *gin.Context) {
	var blogCmt models.BlogCmt
	c.BindJSON(&blogCmt)

	if err := models.CreateBlogCmt(&blogCmt); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "成功评论文章", "blogCmt_info": blogCmt})
	}
}

// GetCommentListOfBlog
// @Summary 获取文章的评论列表
// @Description 根据博客id获取评论列表
// @Tags BlogCmt
// @Accept  json
// @Produce  json
// @Param   blog_id    path    uint     true      "blog_id"
// @Router /v1/blogCmt/getList/{blog_id} [get]
func GetCommentListOfBlog(c *gin.Context) {
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
	cmtList, err := models.GetCmtListOfBlog(uint(blogId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "msg": "成功获取文章评论", "blogCmt_list": cmtList})
}

// DeleteBlogCmt
// @Summary 删除文章评论
// @Description 通过blogCmtID，删除文章评论 eg：{ "blogCmtID":"5"}
// @Tags BlogCmt
// @Accept  json
// @Param   blog_cmt_id    path    uint     true      "blog_cmt_id"
// @Router /v1/blogCmt/delete/{blog_cmt_id} [delete]
func DeleteBlogCmt(c *gin.Context) {
	id, ok := c.Params.Get("blog_cmt_id}")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	blogCmtID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	if err := models.DeleteBlogCmt(uint(blogCmtID)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"msg":  "成功删除文章评论",
		})
	}
}
