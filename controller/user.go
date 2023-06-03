package controller

import (
	"FurballCommunity_backend/config/token"
	"FurballCommunity_backend/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag/example/basic/web"
	"gorm.io/gorm"
	"net/http"
)

// 定义返回状态码
const (
	reStatusError      = 0 //返回常量为0，发生错误
	reStatusSuccess    = 1 //返回常量为1，成功
	reStatusNameRepeat = 2 //返回常量为2，注册用户名重复
)

// Register 注册
// @Summary 用户注册
// @Description 注册一个新的用户 eg：{ "account":"wbq", "password":"123" }
// @Accept  json
// @Produce  json
// @Param   user    body    string   true      "account+password"
// @Success 200 {string} string	"ok"
// @Router /v1/user/register [post]
func Register(c *gin.Context) {
	// 1、从请求中读取数据
	var user models.User
	c.BindJSON(&user)

	//2、判断用户是否存在
	_, err := models.GetUserByAccount(user.Account)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		//注册用户名已存在，输出状态2
		c.JSON(http.StatusOK, gin.H{"state": reStatusError, "text": "此用户已存在！"})
	} else {
		fmt.Println(user.Account, user.Password)
		// 3、存入数据库
		if err := models.CreateUser(&user); err != nil {
			c.JSON(http.StatusCreated, gin.H{"state": reStatusError, "text": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"state": reStatusSuccess, "text": "注册成功！", "userid": user.UserID})
		}
	}
}

// Login 登录
// @Summary 用户登录
// @Description 通过id和pw登录 eg：{ "account":"wbq", "password":"123" }
// @Accept  json
// @Produce  json
// @Param   user    body    string     true      "account+password"
// @Success 200 {string} string	"ok"
// @Router /v1/user/login [post]
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
				ID:       query_user.UserID,
				Username: query_user.Username,
				Account:  query_user.Account,
			})
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"state": reStatusError, "text": err})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"user_id":  query_user.UserID,
				"username": query_user.Username,
				"state":    reStatusSuccess,
				"token":    token,
				"text":     "登陆成功！",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"state": reStatusError, "text": "密码错误！"})
		}
	} else {
		// 用户不存在
		c.JSON(http.StatusOK, gin.H{"state": reStatusError, "text": "登录失败！此用户尚未注册！"})
	}
}

// GetUserList
// @Summary 获取用户列表
// @Description 获取所有用户信息
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
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

// UpdateUserName
// @Summary 更改用户名
// @Description 通过id，修改用户名 eg：{"username":"wangwang" }
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "id"
// @Param   user    body    string     true      "new_name"
// @Success 200 {string} string	"ok"
// @Router /v1/user/updateUsername/{id} [put]
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

// UpdatePassword
// @Summary 更改密码
// @Description 通过id，修改密码 eg：{"password":"123" }
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "id"
// @Param   user    body    string     true      "new_pwd"
// @Router /v1/user/updatePassword/{id} [put]
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

// DeleteUser
// @Summary 删除用户
// @Description 通过id，删除用户 eg：{ "id":"7"}
// @Accept  json
// @Param   id    path    uint     true      "id"
// @Router /v1/user/deleteUser/{id} [delete]
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

// NotFound 设置默认路由当访问一个错误网站时返回
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 ,url not exists!",
	})
}
