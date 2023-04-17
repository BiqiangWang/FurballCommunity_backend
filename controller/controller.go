package controller

import (
	"fmt"
	"net/http"

	"FurballCommunity_backend/config/token"
	"FurballCommunity_backend/models"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 定义response返回的状态码
const (
	resStatusError      = 0 //返回常量为0，发生错误
	resStatusSuccess    = 1 //返回常量为1，成功
	resStatusNameRepeat = 2 //返回常量为2，注册用户名重复
)

// 注册
func Register(c *gin.Context) {
	// 1、从请求中读取数据
	var user models.User
	c.BindJSON(&user)

	//2、判断用户是否存在
	_, err := models.GetUserByAccount(user.Account)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		//注册用户名已存在，输出状态2
		c.JSON(http.StatusOK, gin.H{"state": resStatusError, "text": "此用户已存在！"})
	} else {
		fmt.Println(user.Account, user.Password)
		// 3、存入数据库
		if err := models.CreateUser(&user); err != nil {
			c.JSON(http.StatusCreated, gin.H{"state": resStatusError, "text": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"state": resStatusSuccess, "text": "注册成功！"})
		}
	}
}

// 登录
func Login(c *gin.Context) {
	// 1、从请求中读取数据
	var user models.User
	c.BindJSON(&user)

	//2、先判断用户是否存在
	query_user, err := models.GetUserByAccount(user.Account)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		//3、存在再判断密码是否正确
		if query_user.Password == user.Password {
			// 生成token
			token, err := token.CreateToken(token.UserInfo{
				ID:       query_user.ID,
				Username: query_user.Username,
				Account:  query_user.Account,
			})
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"state": resStatusError, "text": err})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"user_id":  query_user.ID,
				"username": query_user.Username,
				"state":    resStatusSuccess,
				"token":    token,
				"text":     "登陆成功！",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"state": resStatusError, "text": "密码错误！"})
		}
	} else {
		// 用户不存在
		c.JSON(http.StatusOK, gin.H{"state": resStatusError, "text": "登录失败！此用户尚未注册！"})
	}
}

// @Router /v1/user/getUserList [get]
func GetUserList(c *gin.Context) {
	userList, err := models.GetUserList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"list": userList,
		})
	}
}

func UpdateUserName(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id！"})
		return
	}
	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&user)

	if err := models.UpdateUserName(user); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户名修改成功！",
		})
	}
}

func UpdatePassword(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id！"})
		return
	}
	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&user)

	if err := models.UpdatePassword(user); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "密码修改成功！",
		})
	}
}

func DeleteUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteUser(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除成功",
		})
	}
}

// 设置默认路由当访问一个错误网站时返回
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 ,page not exists!",
	})
}
