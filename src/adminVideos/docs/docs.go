// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-03-08 21:36:22.0261359 +0800 CST m=+0.293834201

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "这是服务器演示",
        "title": "这是swagger文档",
        "contact": {
            "name": "何智民",
            "url": "http://127.0.0.1:8000",
            "email": "891453178@qq.com"
        },
        "license": {
            "name": "项目1.0",
            "url": "http://127.0.0.1:8000"
        },
        "version": "1.0"
    },
    "host": "127.0.0.1:8000",
    "basePath": "/api/v1",
    "paths": {}
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
