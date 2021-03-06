// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-08-25 17:09:23.969775 -0700 PDT m=+0.068611435

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Provide a service endpoint that will generate a unique URI",
        "title": "Sample Golang RESTFul APIs",
        "contact": {
            "name": "Arghanil",
            "email": "ARghanil@gmail.com"
        },
        "license": {},
        "version": "0.1"
    },
    "host": "http://localhost:8080/",
    "paths": {
        "/": {
            "get": {
                "description": "Unique Endpoint to get the last submitted payload",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "retrieve"
                ],
                "summary": "Unique Endpoint for GET",
                "operationId": "GetLastContent",
                "responses": {
                    "302": {
                        "description": "Found"
                    },
                    "404": {
                        "description": "Endpoint not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "post": {
                "description": "Unique Endpoint to submit a payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "submit"
                ],
                "summary": "Unique Endpoint for POST",
                "operationId": "SubmitContent",
                "responses": {
                    "202": {
                        "description": "submitted"
                    },
                    "400": {
                        "description": "Bad request body"
                    },
                    "404": {
                        "description": "Endpoint not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/generate": {
            "put": {
                "description": "Create an Endpoint which is unique and accepts GET/POST",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "generate"
                ],
                "summary": "Create a Unique Endpoint",
                "operationId": "GenerateUniqueEndpoint",
                "responses": {
                    "201": {
                        "description": "Endpoint created: url is returned",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.EndpointOutput"
                        }
                    },
                    "500": {
                        "description": "There was as error during processing your request. "
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Unique Endpoint healthcheck",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Unique Endpoint health",
                "operationId": "Health",
                "responses": {
                    "200": {
                        "description": "Status:healthy"
                    },
                    "404": {
                        "description": "Status:unhealthy"
                    },
                    "500": {
                        "description": "Status:unhealthy"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.EndpointOutput": {
            "type": "object",
            "properties": {
                "appurl": {
                    "type": "string"
                },
                "note": {
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
