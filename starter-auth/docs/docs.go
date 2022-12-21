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
        "/auth/permission-create": {
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
                    "permission"
                ],
                "summary": "创建资源权限许可",
                "parameters": [
                    {
                        "description": "创建资源权限许可",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.PermissionReqVO"
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
                                            "$ref": "#/definitions/vo.PermissionRspVO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/role-create": {
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
                    "role"
                ],
                "summary": "创建角色，需要租户管理员权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "租户ID",
                        "name": "Tenant-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建角色",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.RoleCreateReqVO"
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
                                            "$ref": "#/definitions/vo.RoleCreateRspVO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/role-permission-assignment": {
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
                    "role"
                ],
                "summary": "角色赋予资源许可，需要租户管理员权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "租户ID",
                        "name": "Tenant-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "角色赋予资源许可",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.RolePermissionAssignmentReqVO"
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
                                            "$ref": "#/definitions/vo.RolePermissionAssignmentRspVO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/sadmin-token-create": {
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
                "summary": "获得超管Token",
                "parameters": [
                    {
                        "description": "获得超管Token",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.SadminTokenCreateReqVO"
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
                                            "$ref": "#/definitions/vo.SadminTokenCreateRspVO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/tenant-create": {
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
        "/auth/tenant-token-create": {
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
        "/auth/user-authorization-code": {
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
                "summary": "登录验证用户账号密码，验证通过后授权码60秒有效，关联用户信息。可以用授权码接口换取登录Token。",
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
        "/auth/user-create": {
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
        },
        "/auth/user-role-assignment": {
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
                "summary": "用户赋予角色，需要租户管理员权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "租户ID",
                        "name": "Tenant-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "用户赋予角色",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.UserRoleAssignmentReqVO"
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
                                            "$ref": "#/definitions/vo.UserRoleAssignmentRspVO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/user-token-by-authcode": {
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
                "summary": "使用授权码获取用户token信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "租户ID",
                        "name": "Tenant-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "使用授权码获取用户token信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.UserTokenByAuthcodeReqVO"
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
                                            "$ref": "#/definitions/vo.UserTokenByAuthcodeRspVO"
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
        "vo.PermissionInfoVO": {
            "type": "object",
            "required": [
                "authCode",
                "permissionName"
            ],
            "properties": {
                "authCode": {
                    "description": "资源编码",
                    "type": "string"
                },
                "permissionName": {
                    "description": "许可名称",
                    "type": "string"
                }
            }
        },
        "vo.PermissionReqVO": {
            "type": "object",
            "required": [
                "permissions"
            ],
            "properties": {
                "permissions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.PermissionInfoVO"
                    }
                }
            }
        },
        "vo.PermissionRspVO": {
            "type": "object",
            "properties": {
                "permissionIDs": {
                    "description": "许可ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "vo.RoleCreateReqVO": {
            "type": "object",
            "properties": {
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.RoleInfoVO"
                    }
                }
            }
        },
        "vo.RoleCreateRspVO": {
            "type": "object",
            "properties": {
                "roleIDs": {
                    "description": "角色ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "vo.RoleInfoVO": {
            "type": "object",
            "required": [
                "roleCode",
                "roleName"
            ],
            "properties": {
                "roleCode": {
                    "description": "角色编码",
                    "type": "string"
                },
                "roleName": {
                    "description": "角色名",
                    "type": "string"
                }
            }
        },
        "vo.RolePermissionAssignmentReqVO": {
            "type": "object",
            "required": [
                "authCodes"
            ],
            "properties": {
                "authCodes": {
                    "description": "资源编码",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "roleCode": {
                    "description": "角色编码，与角色ID二选一",
                    "type": "string"
                },
                "roleID": {
                    "description": "角色ID，与角色编码二选一",
                    "type": "integer"
                }
            }
        },
        "vo.RolePermissionAssignmentRspVO": {
            "type": "object",
            "properties": {
                "roleID": {
                    "description": "角色ID",
                    "type": "integer"
                }
            }
        },
        "vo.SadminTokenCreateReqVO": {
            "type": "object",
            "required": [
                "secretKey"
            ],
            "properties": {
                "secretKey": {
                    "description": "租户密钥",
                    "type": "string"
                }
            }
        },
        "vo.SadminTokenCreateRspVO": {
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
        "vo.TenantCreateReqVO": {
            "type": "object",
            "properties": {
                "tenants": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.TenantInfoVO"
                    }
                }
            }
        },
        "vo.TenantCreateRspVO": {
            "type": "object",
            "properties": {
                "secretKeys": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.TenantSecretKeyVO"
                    }
                }
            }
        },
        "vo.TenantInfoVO": {
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
        "vo.TenantSecretKeyVO": {
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
                "users"
            ],
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.UserInfoVO"
                    }
                }
            }
        },
        "vo.UserCreateRspVO": {
            "type": "object",
            "properties": {
                "userIDs": {
                    "description": "用户ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "vo.UserInfoVO": {
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
                "userName": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "vo.UserRoleAssignmentReqVO": {
            "type": "object",
            "required": [
                "roleCodes"
            ],
            "properties": {
                "roleCodes": {
                    "description": "角色编码",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "userID": {
                    "description": "用户ID，与用户名二选一",
                    "type": "integer"
                },
                "userName": {
                    "description": "用户名，与用户ID二选一",
                    "type": "string"
                }
            }
        },
        "vo.UserRoleAssignmentRspVO": {
            "type": "object",
            "properties": {
                "userID": {
                    "description": "用户ID",
                    "type": "integer"
                }
            }
        },
        "vo.UserTokenByAuthcodeReqVO": {
            "type": "object",
            "required": [
                "authorizationCode"
            ],
            "properties": {
                "authorizationCode": {
                    "description": "授权码",
                    "type": "string"
                }
            }
        },
        "vo.UserTokenByAuthcodeRspVO": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "clover-auth-api",
	Description:      "clover-auth接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
