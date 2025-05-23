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
        "/tasks": {
            "get": {
                "description": "Возвращает список всех задач.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Получить все задачи",
                "responses": {
                    "200": {
                        "description": "Список задач",
                        "schema": {
                            "$ref": "#/definitions/app.GetAllTasksResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            },
            "post": {
                "description": "Создает новую задачу с указанным заголовком и описанием.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Создать новую задачу",
                "parameters": [
                    {
                        "description": "Данные для создания задачи",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Задача успешно создана",
                        "schema": {
                            "$ref": "#/definitions/app.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Неверный формат запроса"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/tasks/{id}": {
            "put": {
                "description": "Обновляет задачу с указанным ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Обновить задачу",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID задачи",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления задачи",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Задача успешно обновлена"
                    },
                    "400": {
                        "description": "Неверный формат запроса"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            },
            "delete": {
                "description": "Удаляет задачу с указанным ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Удалить задачу",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID задачи",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Задача успешно удалена"
                    },
                    "400": {
                        "description": "Неверный формат ID задачи"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        }
    },
    "definitions": {
        "api_model.TaskResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "app.CreateRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "app.CreateResponse": {
            "type": "object",
            "properties": {
                "task": {
                    "$ref": "#/definitions/api_model.TaskResponse"
                }
            }
        },
        "app.GetAllTasksResponse": {
            "type": "object",
            "properties": {
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api_model.TaskResponse"
                    }
                }
            }
        },
        "app.UpdateRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
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
	Title:            "Todo API",
	Description:      "API для управления задачами",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
