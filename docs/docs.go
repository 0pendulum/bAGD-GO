// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2022-03-03 18:53:59.4041375 +0800 CST m=+0.053718801

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/0pendulum/bAGD-GO",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/good": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增商品",
                "parameters": [
                    {
                        "minLength": 1,
                        "description": "品牌名称",
                        "name": "brand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "minLength": 1,
                        "description": "产地分布图",
                        "name": "map_production",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "配料比例图",
                        "name": "map_Ingredient",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "品牌销量趋势图",
                        "name": "map_trends",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "能量值",
                        "name": "energy",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "蛋白质",
                        "name": "protein",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "脂肪",
                        "name": "fat",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "碳水化合物",
                        "name": "carbohydrates",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "矿物质",
                        "name": "minerals",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "其他",
                        "name": "other",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Good"
                        }
                    },
                    "400": {
                        "description": "请求错误"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Good": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "carbohydrates": {
                    "description": "碳水化合物",
                    "type": "integer"
                },
                "energy": {
                    "description": "能量",
                    "type": "integer"
                },
                "fat": {
                    "description": "脂肪",
                    "type": "integer"
                },
                "mapIngredient": {
                    "description": "配料表",
                    "type": "string"
                },
                "mapProduction": {
                    "description": "产地图",
                    "type": "string"
                },
                "mapTrends": {
                    "description": "品牌趋势",
                    "type": "string"
                },
                "minerals": {
                    "description": "矿物质",
                    "type": "integer"
                },
                "other": {
                    "description": "其他",
                    "type": "integer"
                },
                "protein": {
                    "description": "蛋白质",
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "基于AR的商品信息展示系统",
	Description: "基于AR的商品信息展示系统",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
