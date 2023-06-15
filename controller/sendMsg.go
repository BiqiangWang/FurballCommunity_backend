// 发送验证码功能
package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
	"github.com/gin-gonic/gin"

	"FurballCommunity_backend/utils/redis"
)

type phoneStruct struct {
	Phone string `json:"phone"`
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func createClient() (_result *dysmsapi20170525.Client, _err error) {
	// 初始化Credential。
	credential, _err := credentials.NewCredential(nil)
	if _err != nil {
		panic(_err)
	}
	config := &openapi.Config{
		// 为防止AccessKey ID、Secret泄露，这里使用Credential配置凭证。
		// Credential配置凭证：ID、Secret存放在本机的系统环境变量中，需手动添加环境变量
		Credential: credential,

		// 您的 AccessKey ID
		// AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		// AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	// _result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendMsg(c *gin.Context) {
	var phone phoneStruct
	c.BindJSON(&phone)

	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。
	// 此处采用的Credential配置凭证，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := createClient()
	if _err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": _err.Error()})
		return
	}
	// 生成6位数字随机验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("毛球社区"),
		TemplateCode:  tea.String("SMS_460725820"),
		PhoneNumbers:  tea.String(phone.Phone),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": _err.Error()})
			return
		}
		// 将手机号与验证码保存下来，5分钟有效
		err := redis.RedisSet(phone.Phone, code, 5*time.Minute)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success"})
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": _err.Error()})
			return
		}
	}
}
