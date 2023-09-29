// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API AdoPet",
            "url": "https://www.google.com/",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/abrigo": {
            "get": {
                "description": "Rota para consultar todos os Abrigos",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Abrigo"
                ],
                "summary": "Consulta todos os Abrigos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Abrigo"
                        }
                    }
                }
            },
            "post": {
                "description": "Criar um novo perfil de Abrigo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Abrigo"
                ],
                "summary": "Criar um novo perfil de Abrigo",
                "parameters": [
                    {
                        "description": "Informe os dados do Abrigo",
                        "name": "Abrigo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Abrigo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Abrigo"
                        }
                    }
                }
            }
        },
        "/abrigo/{id}": {
            "put": {
                "description": "Alterar um abrigo no banco de dados",
                "produces": [
                    "application/json"
                ],
                "summary": "Alterar um abrigo existente.",
                "operationId": "alterAbrigo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do abrigo a ser alterado",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dados do abrigo a serem atualizado",
                        "name": "abrigo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Abrigo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Abrigo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/home": {
            "get": {
                "description": "Rota para testar a API",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "home"
                ],
                "summary": "Rota para testar a API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tutor": {
            "get": {
                "description": "Rota para consultar todos os Tutores",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tutor"
                ],
                "summary": "Consulta todos os Tutores",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tutor"
                        }
                    }
                }
            },
            "post": {
                "description": "Criar um novo Tutor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tutor"
                ],
                "summary": "Criar um novo Tutor",
                "parameters": [
                    {
                        "description": "Informe os dados do Tutor",
                        "name": "Tutor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Tutor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tutor"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Abrigo": {
            "type": "object",
            "properties": {
                "cidade": {
                    "type": "string"
                },
                "foto": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "sobre": {
                    "type": "string"
                },
                "telefone": {
                    "type": "string"
                },
                "uf": {
                    "type": "string"
                }
            }
        },
        "models.Tutor": {
            "type": "object",
            "properties": {
                "cidade": {
                    "type": "string"
                },
                "foto": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "sobre": {
                    "type": "string"
                },
                "telefone": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "AdoPet",
	Description:      "API para o aplicativo AdoPet",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
