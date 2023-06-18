package controller

import (
	"FurballCommunity_backend/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag/example/basic/web"
	"gorm.io/gorm"
)

// 定义返回状态码
const (
	reStatusError   = 0 //返回常量为0，发生错误
	reStatusSuccess = 1 //返回常量为1，成功
)

// Register 注册
// @Summary 用户注册
// @Description 注册一个新的用户 eg：{ "account":"wbq", "password":"123" }
// @Tags User
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
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "此用户账号已存在！"})
	} else {
		// 3、存入数据库
		if err := models.CreateUser(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		} else {
			// 生成token
			genToken, err := token.CreateToken(token.UserInfo{
				ID:       user.UserID,
				Username: user.Username,
				Account:  user.Account,
			})
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err})
				return
			}
			// 将token存到redis
			err = token.SetTokenCache(user.UserID, genToken)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err})
				return
			}
			c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "token": genToken, "msg": "注册成功！", "user": user})
		}
	}
}

// Login 账号密码登录
// @Summary 用户登录
// @Description 通过id和pw登录 eg：{ "account":"wbq", "password":"123" }
// @Tags User
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
			genToken, err := token.CreateToken(token.UserInfo{
				ID:       query_user.UserID,
				Username: query_user.Username,
				Account:  query_user.Account,
			})
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err})
				return
			}
			// 将token存到redis
			err = token.SetTokenCache(query_user.UserID, genToken)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"user":  query_user,
				"code":  reStatusSuccess,
				"token": genToken,
				"msg":   "登陆成功！",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "密码错误！"})
		}
	} else {
		// 用户不存在
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "登录失败！此用户尚未注册！"})
	}
}

type phoneParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// LoginWithPhone
// @Summary 手机验证码登录
// @Description 通过id和pw登录 eg：{ "phone":"13533337492", "code":"123456" }
// @Tags User
// @Accept  json
// @Produce  json
// @Param   user    body    string     true      "phone+code"
// @Success 200 {string} string	"ok"
// @Router /v1/user/loginWithPhone [post]
func LoginWithPhone(c *gin.Context) {
	// 1、从请求中读取数据
	var param phoneParam
	c.BindJSON(&param)
	var user models.User
	user.Phone = param.Phone
	//2、先判断用户是否存在
	query_user, err := models.GetUserByPhone(user.Phone)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		//3、存在再判断短信验证码是否正确
		redis_code, err := redis.RedisGet(user.Phone)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err})
			return
		}
		if param.Code == redis_code {
			// 生成token
			genToken, err := token.CreateToken(token.UserInfo{
				ID:       query_user.UserID,
				Username: query_user.Username,
				Account:  query_user.Account,
			})
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err})
				return
			}
			redis.RedisSet(user.Phone, "", 0) //将手机号对应的短信验证码的redis缓存设为""
			// 将token存到redis
			err = token.SetTokenCache(query_user.UserID, genToken)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"user":  query_user,
				"code":  reStatusSuccess,
				"token": genToken,
				"msg":   "登陆成功！",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "验证码错误或已失效！"})
		}
	} else {
		// 用户不存在,则直接注册,初始账号、密码、用户名均设置为手机号
		// 3、存入数据库
		user.Account = user.Phone
		user.Password = user.Phone
		user.Username = user.Phone
		if err := models.CreateUser(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		} else {
			redis.RedisSet(user.Phone, "", 0) //将手机号对应的短信验证码的redis缓存设为""
			// 生成token
			genToken, err := token.CreateToken(token.UserInfo{
				ID:       query_user.UserID,
				Username: query_user.Username,
				Account:  query_user.Account,
			})
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err})
				return
			}
			// 将token存到redis
			err = token.SetTokenCache(query_user.UserID, genToken)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err})
				return
			}
			c.JSON(http.StatusOK, gin.H{"code": reStatusSuccess, "token": genToken, "msg": "注册成功！", "user": user})
		}
	}
}

// GetUserList
// @Summary 获取用户列表
// @Description 获取所有用户信息
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /v1/user/getUserList [get]
func GetUserList(c *gin.Context) {
	userList, err := models.GetUserList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"list": userList,
		})
	}
}

// UpdateUserName
// @Summary 更改用户名
// @Description 通过id，修改用户名 eg：{"username":"wangwang" }
// @Tags User
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "id"
// @Param   user    body    string     true      "new_name"
// @Success 200 {string} string	"ok"
// @Router /v1/user/updateUsername/{id} [put]
func UpdateUserName(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "无效的id！"})
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	user, err := models.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.BindJSON(&user)
	if err := models.UpdateUserName(user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户名修改成功！",
		})
	}
}

// UpdatePassword
// @Summary 更改密码
// @Description 通过id，修改密码 eg：{"password":"123" }
// @Tags User
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "id"
// @Param   user    body    string     true      "new_pwd"
// @Router /v1/user/updatePassword/{id} [put]
func UpdatePassword(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "无效的id！"})
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	user, err := models.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.BindJSON(&user)

	if err := models.UpdatePassword(user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "密码修改成功！",
		})
	}
}

// DeleteUser
// @Summary 删除用户
// @Description 通过id，删除用户 eg：{ "id":"7"}
// @Tags User
// @Accept  json
// @Param   id    path    uint     true      "id"
// @Router /v1/user/deleteUser/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "无效的id"})
		return
	}
	if err := models.DeleteUser(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "删除成功",
		})
	}
}

// GetUserInfo
// @Summary 获取用户信息
// @Description 根据用户id获取用户信息
// @Tags User
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "id"
// @Router /v2/user/getUserById/{id} [get]
func GetUserInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	if user, err := models.GetUserById(uint(userId)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"msg":  "成功获取用户信息",
			"user": user,
		})
	}
}

// UpdateUserInfo
// @Summary 更改用户信息
// @Description 通过id，更新用户信息，包括手机号、权限等级、性别、地址、分数、简介、身份证号、头像、养宠经验、工作时间、可养宠数量和身份证照片等，用户名和密码由原接口修改 eg：{"gender":1}
// @Tags User
// @Accept  json
// @Produce  json
// @Param   id    path    uint     true      "id"
// @Param   user    body    string     true      "new_pet_info"
// @Success 200 {string} string	"ok"
// @Router /v2/user/updateUserInfo/{id} [put]
func UpdateUserInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": "无效的id"})
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": reStatusError, "msg": "转换后无效的id"})
		return
	}
	user, err := models.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
		return
	}
	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": e.Error()})
		return
	}
	if err := models.UpdateUserInfo(user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": reStatusError, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": reStatusSuccess,
			"meg":  "成功修改用户信息！",
			"info": user,
		})
	}

}

// NotFound 设置默认路由当访问一个错误网站时返回
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"msg":    "404 ,url not exists!",
	})
}
