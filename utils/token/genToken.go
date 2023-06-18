package token

import (
	"FurballCommunity_backend/utils"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 自定义jwt的声明字段信息+标准字段
type CustomClaims struct {
	UserInfo UserInfo
	jwt.StandardClaims
}

type UserInfo struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Account  string `json:"account"  binding:"required"`
	Username string `json:"username"`
}

var mySigningKey = []byte("woshisuperadminfangguowoba")

// 生成token
func CreateToken(userInfo UserInfo) (string, error) {
	d := CustomClaims{
		UserInfo: userInfo,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,     //生效时间
			ExpiresAt: time.Now().Unix() + 3600*8, //过期时间 8小时
			Issuer:    "admin",                    //签发人
		},
	}
	fmt.Println(d.ExpiresAt)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, d)
	s, err := t.SignedString(mySigningKey) //对token进行加密
	if err != nil {
		log.Println(err)
	}
	return s, err
}

// 解析Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if token == nil {
		return nil, errors.New(utils.ErrorsTokenInvalid)
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(utils.ErrorsTokenMalFormed)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(utils.ErrorsTokenNotActiveYet)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(utils.ErrorTokenExpired)
			} else {
				return nil, errors.New(utils.ErrorsTokenInvalid)
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		log.Println("claims:", claims)
		return claims, nil
	} else {
		return nil, errors.New(utils.ErrorsTokenInvalid)
	}
}

// 更新token
func RefreshToken(tokenString string) (string, error) {
	if CustomClaims, err := ParseToken(tokenString); err == nil {
		// CustomClaims.ExpiresAt = time.Now().Unix() + 60*60*2
		return CreateToken(CustomClaims.UserInfo)
	} else {
		return "", err
	}
}
