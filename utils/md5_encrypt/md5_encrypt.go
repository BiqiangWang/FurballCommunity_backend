package md5_encrypt

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func MD5(params string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(params))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// 如果需要对用户密码加密存储，则可用此方法
// 先base64，然后MD5
func Base64Md5(params string) string {
	return MD5(base64.StdEncoding.EncodeToString([]byte(params)))
}
