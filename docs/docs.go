// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/admin/category-create": {
            "post": {
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "分类创建",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "parentId",
                        "name": "parentId",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/category-delete": {
            "delete": {
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "分类删除",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/category-list": {
            "get": {
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "分类列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/category-modify": {
            "put": {
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "分类修改",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "parentId",
                        "name": "parentId",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/problem-create": {
            "post": {
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "问题创建",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "max_runtime",
                        "name": "max_runtime",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "max_mem",
                        "name": "max_mem",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "category_ids",
                        "name": "category_ids",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "test_cases",
                        "name": "test_cases",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/problem-modify": {
            "put": {
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "问题修改",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "max_runtime",
                        "name": "max_runtime",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "max_mem",
                        "name": "max_mem",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "category_ids",
                        "name": "category_ids",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "test_cases",
                        "name": "test_cases",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "tags": [
                    "公共方法"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/problem-detail": {
            "get": {
                "tags": [
                    "公共方法"
                ],
                "summary": "问题详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "problem identity",
                        "name": "identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/problem-list": {
            "get": {
                "tags": [
                    "公共方法"
                ],
                "summary": "问题列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "category_identity",
                        "name": "category_identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rank-list": {
            "get": {
                "tags": [
                    "公共方法"
                ],
                "summary": "用户排行榜",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "tags": [
                    "公共方法"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "mail",
                        "name": "mail",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "code",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "phone",
                        "name": "phone",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/send-code": {
            "post": {
                "tags": [
                    "公共方法"
                ],
                "summary": "发送验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/submit-list": {
            "get": {
                "tags": [
                    "公共方法"
                ],
                "summary": "提交列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "problem_identity",
                        "name": "problem_identity",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "user_identity",
                        "name": "user_identity",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user-detail": {
            "get": {
                "tags": [
                    "公共方法"
                ],
                "summary": "用户详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "problem identity",
                        "name": "identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/submit": {
            "post": {
                "tags": [
                    "用户私有方法"
                ],
                "summary": "代码提交",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "problem_identity",
                        "name": "problem_identity",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "code",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
