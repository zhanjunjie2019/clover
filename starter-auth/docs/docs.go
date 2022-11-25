// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/tenant-create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "创建租户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "租户ID",
                        "name": "Tenant-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建租户",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.TenantCreateReqVO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/vo.TenantCreateRspVO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/tenant-token-create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "创建租户Token",
                "parameters": [
                    {
                        "description": "创建租户Token",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.TenantTokenCreateReqVO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/vo.TenantTokenCreateRspVO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user-authorization-code": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "登录验证用户账号密码，验证通过后在Redis保存一个授权码60秒有效，关联用户信息。用以可以用授权码接口换取登录Token。",
                "parameters": [
                    {
                        "type": "string",
                        "description": "租户ID",
                        "name": "Tenant-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "用户登录获得授权码，注意授权码不是Token，不能直接用于访问接口",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.UserAuthorizationCodeReqVO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/vo.UserAuthorizationCodeRspVO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user-create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "创建用户，需要租户管理员权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "租户ID",
                        "name": "Tenant-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建用户",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.UserCreateReqVO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/vo.UserCreateRspVO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误码：成功[0],底层的错误[1000-1999],平台层级的错误[2000-2999]",
                    "type": "integer"
                },
                "data": {
                    "description": "数据"
                },
                "msg": {
                    "description": "详细信息",
                    "type": "string"
                }
            }
        },
        "vo.TenantCreateReqVO": {
            "type": "object",
            "required": [
                "tenantName"
            ],
            "properties": {
                "AccessTokenTimeLimit": {
                    "description": "访问Token有效时限，非必要，默认7200s",
                    "type": "integer"
                },
                "redirectUrl": {
                    "description": "授权码重定向路径，非必要",
                    "type": "string"
                },
                "tenantID": {
                    "description": "租户ID，非必要，不传默认则随机生成",
                    "type": "string"
                },
                "tenantName": {
                    "description": "租户名",
                    "type": "string"
                }
            }
        },
        "vo.TenantCreateRspVO": {
            "type": "object",
            "properties": {
                "secretKey": {
                    "description": "租户密钥",
                    "type": "string"
                },
                "tenantID": {
                    "description": "租户ID",
                    "type": "string"
                }
            }
        },
        "vo.TenantTokenCreateReqVO": {
            "type": "object",
            "required": [
                "secretKey",
                "tenantID"
            ],
            "properties": {
                "accessTokenExpirationTime": {
                    "description": "访问Token过期时间戳，非必要，不传则按当前时间+戳追加租户设置有效时限",
                    "type": "integer"
                },
                "secretKey": {
                    "description": "租户密钥",
                    "type": "string"
                },
                "tenantID": {
                    "description": "租户ID",
                    "type": "string"
                }
            }
        },
        "vo.TenantTokenCreateRspVO": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "description": "访问Token",
                    "type": "string"
                },
                "accessTokenExpirationTime": {
                    "description": "访问Token过期时间戳",
                    "type": "integer"
                }
            }
        },
        "vo.UserAuthorizationCodeReqVO": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "redirectUrl": {
                    "description": "重定向路径，非必要，不传则按租户设置",
                    "type": "string"
                },
                "userName": {
                    "description": "账户名",
                    "type": "string"
                }
            }
        },
        "vo.UserAuthorizationCodeRspVO": {
            "type": "object",
            "properties": {
                "authorizationCode": {
                    "description": "授权码",
                    "type": "string"
                },
                "redirectUrl": {
                    "description": "重定向路径",
                    "type": "string"
                }
            }
        },
        "vo.UserCreateReqVO": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "password": {
                    "description": "Password 密码",
                    "type": "string"
                },
                "userName": {
                    "description": "UserName 用户名",
                    "type": "string"
                }
            }
        },
        "vo.UserCreateRspVO": {
            "type": "object",
            "properties": {
                "userId": {
                    "description": "UserId 用户ID",
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "C-Token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1.0.0",
	Host:             "",
	BasePath:         "/auth",
	Schemes:          []string{},
	Title:            "clover-auth-api",
	Description:      "clover-auth接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
