// 图形验证码模块
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type verify struct {
	CaptchaId   string `json:"captchaId"`
	VerifyValue string `json:"verifyValue"`
}

// configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// GenerateCaptchaHandler base64Captcha create http handler
// @Summary 获取图形验证码
// @Description 获取一张图形验证码，同时返回captchaId
// @Tags captcha
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /v1/api/getCaptcha [get]
func GenerateCaptchaHandler(c *gin.Context) {
	var param configJsonBody = configJsonBody{
		Id:          "",
		CaptchaType: "string",
		VerifyValue: "",
		DriverAudio: &base64Captcha.DriverAudio{},
		DriverString: &base64Captcha.DriverString{
			Length:          4,
			Height:          60,
			Width:           240,
			ShowLineOptions: 2,
			NoiseCount:      0,
			Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		},
		DriverChinese: &base64Captcha.DriverChinese{},
		DriverMath:    &base64Captcha.DriverMath{},
		DriverDigit:   &base64Captcha.DriverDigit{},
	}

	c1 := base64Captcha.NewCaptcha(param.DriverString.ConvertFonts(), store)
	id, b64s, err := c1.Generate()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": b64s, "captchaId": id, "msg": "success"})
}

// CaptchaVerifyHandle 验证图形验证码 base64Captcha verify http handler
/** 传入参数：{
    "CaptchaId":"mFXBu7EueGbtNqsErYdm",
    "VerifyValue":"vvsz"
}
**/
// @Summary 验证图形验证码
// @Description 验证图形验证码 eg：{ "CaptchaId":"mFXBu7EueGbtNqsErYdm", "VerifyValue":"vvsz" }
// @Tags captcha
// @Accept  json
// @Produce  json
// @Param   CaptchaVerifyHandle    body    string   true      "CaptchaId+VerifyValue"
// @Success 200 {string} string	"ok"
// @Router /v1/api/verifyCaptcha [post]
func CaptchaVerifyHandle(c *gin.Context) {

	//parse request json body
	var verifyReq verify
	c.BindJSON(&verifyReq)
	//verify the captcha
	if store.Verify(verifyReq.CaptchaId, verifyReq.VerifyValue, true) {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "ok"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "图形验证码错误"})
	}
}
