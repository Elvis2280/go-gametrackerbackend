// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/games": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games"
                ],
                "summary": "Get all Games",
                "operationId": "get-all-games",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Game"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games"
                ],
                "summary": "Create a game",
                "operationId": "create-game",
                "parameters": [
                    {
                        "description": "Tag",
                        "name": "platform",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Game"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Game"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games"
                ],
                "summary": "Delete a game",
                "operationId": "delete-game",
                "parameters": [
                    {
                        "description": "Tag",
                        "name": "platform",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Game"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Game"
                        }
                    }
                }
            }
        },
        "/games/{id}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games"
                ],
                "summary": "Update a game",
                "operationId": "update-game",
                "parameters": [
                    {
                        "description": "Tag",
                        "name": "platform",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Game"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Game"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login into an account",
                "operationId": "login-account",
                "parameters": [
                    {
                        "description": "Tag",
                        "name": "platform",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/platforms": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platforms"
                ],
                "summary": "Get all platforms",
                "operationId": "get-all-platforms",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Platforms"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platforms"
                ],
                "summary": "Create a platform",
                "operationId": "create-platform",
                "parameters": [
                    {
                        "description": "Tag",
                        "name": "platform",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Platforms"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Platforms"
                        }
                    }
                }
            }
        },
        "/platforms/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platforms"
                ],
                "summary": "Delete a platform",
                "operationId": "delete-platform",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Platform ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Platforms"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Create an account",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "Tag",
                        "name": "platform",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserSignup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/tags": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Get all tags",
                "operationId": "get-all-tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tags"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Create a tag",
                "operationId": "create-tag",
                "parameters": [
                    {
                        "description": "Tag",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Tags"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tags"
                        }
                    }
                }
            }
        },
        "/tags/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Delete a tag",
                "operationId": "delete-tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tag ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tags"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Game": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "platforms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Platforms"
                    }
                },
                "status": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Tags"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Platforms": {
            "type": "object",
            "properties": {
                "iconName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Tags": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "games": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Game"
                    }
                },
                "isVerified": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.UserLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.UserSignup": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Game tracker API",
	Description:      "Game tracker app API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
