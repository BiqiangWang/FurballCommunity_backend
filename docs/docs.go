// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/api/getCaptcha": {
            "get": {
                "description": "获取一张图形验证码，同时返回captchaId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "captcha"
                ],
                "summary": "获取图形验证码",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/api/verifyCaptcha": {
            "post": {
                "description": "验证图形验证码 eg：{ \"CaptchaId\":\"mFXBu7EueGbtNqsErYdm\", \"VerifyValue\":\"vvsz\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "captcha"
                ],
                "summary": "验证图形验证码",
                "parameters": [
                    {
                        "description": "CaptchaId+VerifyValue",
                        "name": "CaptchaVerifyHandle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/order/create": {
            "post": {
                "description": "根据用户id，创建一个新的订单 eg：{ \"pet_id\":3, \"announcer_id\":2 }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "创建订单",
                "parameters": [
                    {
                        "description": "petname + userid",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/order/getOrderInfoById/{order_id}": {
            "get": {
                "description": "根据订单id获取详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "获取订单信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order_id",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/order/getOrderList/{user_id}": {
            "get": {
                "description": "根据用户id获取订单列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "获取用户的订单列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/order/getOrderOfPet/{pet_id}": {
            "get": {
                "description": "根据宠物id获取订单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "获取宠物的订单",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "pet_id",
                        "name": "pet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/order/updateOrderInfo/{order_id}": {
            "put": {
                "description": "通过订单id，更新接收者、开始结束时间、地点、健康、订单状态、备注、价格、评价、评分等 eg：{\"receiver_id\":1}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "更改订单信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order_id",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "new_order_info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/orderCmt/create": {
            "post": {
                "description": "根据用户id、订单id，添加一个新的订单评论 eg：{ \"user_id\":1, \"order_id\":2, \"content\":\"宠物照顾的针布戳呢\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderCmt"
                ],
                "summary": "创建订单评论",
                "parameters": [
                    {
                        "description": "userid + orderId + content",
                        "name": "orderCmt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/orderCmt/deleteOrderCmt/{order_cmt_id}": {
            "delete": {
                "description": "通过orderCmtID，删除宠物 eg：{ \"orderCmtID\":\"5\"}",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "OrderCmt"
                ],
                "summary": "删除订单评论",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order_cmt_id",
                        "name": "order_cmt_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/orderCmt/getOrderCmtList/{order_id}": {
            "get": {
                "description": "根据订单id获取评论列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderCmt"
                ],
                "summary": "获取订单的评论列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order_id",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/pet/add": {
            "post": {
                "description": "添加一个新的宠物 eg：{\"pet_name\":\"xiaohuang\",\"user_id\":2 }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pet"
                ],
                "summary": "添加宠物",
                "parameters": [
                    {
                        "description": "petname + userid",
                        "name": "pet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/pet/deletePet/{id}": {
            "delete": {
                "description": "通过id，删除宠物 eg：{ \"id\":\"5\"}",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Pet"
                ],
                "summary": "删除宠物",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/pet/getPetInfoByID/{id}": {
            "get": {
                "description": "通过宠物id查询宠物信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pet"
                ],
                "summary": "通过宠物id查询宠物信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "pet_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/pet/getPetList/{id}": {
            "get": {
                "description": "根据用户id获取宠物列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pet"
                ],
                "summary": "获取宠物列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/pet/updatePetInfo/{id}": {
            "put": {
                "description": "通过id，更新宠物信息，包括宠物名称、年龄、重量、绝育信息、品种和健康情况等 eg：{\"pet_name\":\"wangwang\", \"gender\":1, \"age\":2, \"weight\":33, \"sterilization\":1, \"breed\":\"taidi\", \"health\":\"yes\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pet"
                ],
                "summary": "更改宠物信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "new_pet_info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user//loginWithPhone": {
            "post": {
                "description": "通过id和pw登录 eg：{ \"phone\":\"13533337492\", \"code\":\"123456\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "手机验证码登录",
                "parameters": [
                    {
                        "description": "phone+code",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/deleteUser/{id}": {
            "delete": {
                "description": "通过id，删除用户 eg：{ \"id\":\"7\"}",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/user/getUserList": {
            "get": {
                "description": "获取所有用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户列表",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/login": {
            "post": {
                "description": "通过id和pw登录 eg：{ \"account\":\"wbq\", \"password\":\"123\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "account+password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/register": {
            "post": {
                "description": "注册一个新的用户 eg：{ \"account\":\"wbq\", \"password\":\"123\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "account+password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/updatePassword/{id}": {
            "put": {
                "description": "通过id，修改密码 eg：{\"password\":\"123\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "更改密码",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "new_pwd",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/user/updateUsername/{id}": {
            "put": {
                "description": "通过id，修改用户名 eg：{\"username\":\"wangwang\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "更改用户名",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "new_name",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/user/getUserById/{id}": {
            "get": {
                "description": "根据用户id获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v2/user/updateUserInfo/{id}": {
            "put": {
                "description": "通过id，更新用户信息，包括手机号、权限等级、性别、地址、分数、简介、身份证号、头像、养宠经验、工作时间、可养宠数量和身份证照片等，用户名和密码由原接口修改 eg：{\"gender\":1}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "更改用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "new_pet_info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "毛球社区",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
