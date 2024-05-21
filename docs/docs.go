// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "summary": "User login",
                "description": "Authenticate user with username and password",
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/LoginRequest"
                        },
                        "required": true,
                        "description": "Login request object"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User logged in successfully"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/register": {
            "post": {
                "summary": "Register a new user",
                "description": "Register a new user with username, password, email, etc.",
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/User"
                        },
                        "required": true,
                        "description": "User object to be registered"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User registered successfully"
                    },
                    "400": {
                        "description": "Invalid input"
                    }
                }
            }
        },
        "/refresh-token": {
            "post": {
                "summary": "Refresh JWT token",
                "description": "Refresh JWT token for authenticated user",
                "tags": [
                    "auth"
                ],
                "responses": {
                    "200": {
                        "description": "Token refreshed successfully"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "summary": "List all users",
                "description": "Get a list of all registered users",
                "tags": [
                    "users"
                ],
                "responses": {
                    "200": {
                        "description": "List of users retrieved successfully"
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "summary": "Upload an image",
                "description": "Upload an image file",
                "tags": [
                    "images"
                ],
                "consumes": ["multipart/form-data"],
                "produces": ["application/json"],
                "parameters": [
                    {
                        "name": "file",
                        "in": "formData",
                        "description": "The image file to upload",
                        "required": true,
                        "type": "images"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image uploaded successfully"
                    },
                    "400": {
                        "description": "Invalid input"
                    }
                }
            }
        },
        "/open-image/{filename}": {
            "get": {
                "summary": "Get an image by filename",
                "description": "Retrieve an image by its filename",
                "tags": [
                    "images"
                ],
                "parameters": [
                    {
                        "name": "filename",
                        "in": "path",
                        "description": "Filename of the image to retrieve",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image retrieved successfully"
                    },
                    "404": {
                        "description": "Image not found"
                    }
                }
            }
        }
    },
    "definitions": {
        "User": {
            "type": "object",
            "properties": {
            
                "username": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "format": "email"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                }
            }
        },
        "LoginRequest": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Project-1 API",
	Description:      "This is a sample server for Project-1.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
