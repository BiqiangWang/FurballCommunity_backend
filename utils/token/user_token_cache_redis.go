package token

import (
	"FurballCommunity_backend/utils/md5_encrypt"
	"FurballCommunity_backend/utils/redis"
	"strconv"
	"time"
)

var userTokenKey = "userToken_"

// SetTokenCache 设置缓存，默认8小时有效期
func SetTokenCache(userId uint, token string) error {
	// 存储用户token时转为MD5，下一步比较的时候可以更加快速地比较是否一致
	return redis.RedisSet(userTokenKey+strconv.FormatUint(uint64(userId), 10), md5_encrypt.MD5(token), 8*time.Hour)
}

// TokenCacheIsExists 查询token是否在redis存在
func TokenCacheIsExists(userId uint) (exists bool, err error) {
	res, err := redis.RedisExists(userTokenKey + strconv.FormatUint(uint64(userId), 10))
	if err != nil {
		// 错误处理
		return false, err
	}
	if res == 1 {
		exists = true // key存在
	} else {
		exists = false // key不存在
	}
	return
}

// ClearUserToken 清除某个用户的token缓存，当用户更改密码或者用户被禁用则删除该用户的token缓存
func ClearUserToken(userId uint) error {
	return redis.RedisDel(userTokenKey + strconv.FormatUint(uint64(userId), 10))
}
