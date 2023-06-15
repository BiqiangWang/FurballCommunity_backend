package token

import "errors"

func VerifyToken(userId uint, token string) (bool, error) {
	exists, err := TokenCacheIsExists(userId)
	if err != nil {
		return false, err
	}
	if !exists { //如果redis缓存不存在说明未登录
		return false, errors.New("用户登录信息不合法，请重新登陆")
	}
	claims, err := ParseToken(token)
	if err != nil {
		return false, err
	} else {
		// 验证签发人是否是"admin"
		if claims.StandardClaims.Issuer != "admin" {
			return false, errors.New("token签发人信息不合法，请重新登陆")
		}
		// 验证签发token中的userid与本人id是否对应
		if claims.UserInfo.ID != userId {
			return false, errors.New("用户登录信息与token不一致，请重新登陆")
		}
	}
	return true, nil
}
