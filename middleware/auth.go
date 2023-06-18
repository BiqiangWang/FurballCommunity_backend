package middleware

import (
	"FurballCommunity_backend/utils/token"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type HeaderParams struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
	UserId        uint   `header:"UserId" binding:"required"`
}

// CheckTokenAuth 检查token完整性、有效性中间件
func CheckTokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerParams := HeaderParams{}

		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			//停止当前请求直接返回，防止后续处理函数继续执行
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -1, "msg": err.Error()})
			return
		}
		tokenParam := strings.Split(headerParams.Authorization, " ")
		if len(tokenParam) == 2 && len(tokenParam[1]) >= 20 {
			tokenIsEffective, err := token.VerifyToken(headerParams.UserId, tokenParam[1])
			if tokenIsEffective {
				if customToken, err := token.ParseToken(tokenParam[1]); err == nil {
					// token验证通过，同时将解析出来的对象绑定在请求上下文
					context.Set("userToken", customToken)
				}
				log.Println("token验证通过")
				context.Next()
			} else {
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -1, "msg": err.Error()})
			}
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -1, "msg": "token格式不合法，请重新登陆"})
		}
	}
}

// CheckCasbinAuth casbin检查用户对应的角色权限是否允许访问接口
func CheckCasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// requstUrl := c.Request.URL.Path
		// method := c.Request.Method

		// 模拟请求参数转换后的角色（roleId=2）
		// 主线版本没有深度集成casbin的使用逻辑
		// GinSkeleton-Admin 系统则深度集成了casbin接口权限管控
		// 详细实现参考地址：https://gitee.com/daitougege/gin-skeleton-admin-backend/blob/master/app/http/middleware/authorization/auth.go
		// role := "2" // 这里模拟某个用户的roleId=2

		// 这里将用户的id解析为所拥有的的角色，判断是否具有某个权限即可
		// isPass, err := variable.Enforcer.Enforce(role, requstUrl, method)
		// if err != nil {
		// 	response.ErrorCasbinAuthFail(c, err.Error())
		// 	return
		// } else if !isPass {
		// 	response.ErrorCasbinAuthFail(c, "")
		// 	return
		// } else {
		// 	c.Next()
		// }
	}
}
