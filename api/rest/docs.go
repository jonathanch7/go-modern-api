// Package rest GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package rest

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/businesses": {
            "get": {
                "description": "Search all businesses",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Businesses"
                ],
                "summary": "Search all businesses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtoresponse.BusinessResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtoresponse.BusinessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtoresponse.BusinessResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new business",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Businesses"
                ],
                "summary": "Create a new business",
                "parameters": [
                    {
                        "description": "Business",
                        "name": "Business",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtorequest.BusinessRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtoresponse.BusinessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtoresponse.BusinessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtoresponse.BusinessResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtorequest.BusinessRequest": {
            "type": "object",
            "required": [
                "businessId",
                "description",
                "name",
                "regDate"
            ],
            "properties": {
                "businessId": {
                    "description": "Business ID",
                    "type": "string",
                    "maxLength": 3,
                    "minLength": 1
                },
                "description": {
                    "description": "Detailed description of business",
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 1
                },
                "name": {
                    "description": "Business name",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1
                },
                "publicCode": {
                    "description": "Public code",
                    "type": "string",
                    "maxLength": 36
                },
                "regDate": {
                    "description": "Register date",
                    "type": "string"
                }
            }
        },
        "dtoresponse.BusinessResponse": {
            "type": "object",
            "required": [
                "regDate"
            ],
            "properties": {
                "businessId": {
                    "description": "Business ID",
                    "type": "string",
                    "maxLength": 3
                },
                "description": {
                    "description": "Detailed description of business",
                    "type": "string",
                    "maxLength": 200
                },
                "name": {
                    "description": "Business name",
                    "type": "string",
                    "maxLength": 50
                },
                "publicCode": {
                    "description": "Public code",
                    "type": "string",
                    "maxLength": 36
                },
                "regDate": {
                    "description": "Register date",
                    "type": "string"
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
	Version:     "1.0.0",
	Host:        "localhost:8080",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Modern API",
	Description: "The purpose of this application is be modern API.",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
