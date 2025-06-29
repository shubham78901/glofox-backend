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
        "/bookings": {
            "get": {
                "description": "Retrieves a list of all bookings",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get all bookings",
                "responses": {
                    "200": {
                        "description": "List of bookings",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/responses.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Booking"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new booking for a member to attend a class",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Create a new booking",
                "parameters": [
                    {
                        "description": "Booking information",
                        "name": "booking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BookingInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Booking created successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/responses.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Booking"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/bookings/{id}": {
            "get": {
                "description": "Retrieves a booking by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get booking by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Booking found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/responses.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Booking"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Booking not found",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/classes": {
            "get": {
                "description": "Retrieves a list of all classes, optionally filtered by date",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classes"
                ],
                "summary": "Get all classes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter classes by date (YYYY-MM-DD)",
                        "name": "date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of classes",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/responses.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Class"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid date format",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new fitness class with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classes"
                ],
                "summary": "Create a new class",
                "parameters": [
                    {
                        "description": "Class information",
                        "name": "class",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ClassInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Class created successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/responses.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Class"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/classes/{id}": {
            "get": {
                "description": "Retrieves a class by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classes"
                ],
                "summary": "Get class by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Class ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Class found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/responses.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Class"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Class not found",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Booking": {
            "type": "object",
            "properties": {
                "classId": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.BookingInput": {
            "type": "object",
            "required": [
                "classId",
                "date",
                "name"
            ],
            "properties": {
                "classId": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Class": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "className": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                }
            }
        },
        "models.ClassInput": {
            "type": "object",
            "required": [
                "capacity",
                "className",
                "endDate",
                "startDate"
            ],
            "properties": {
                "capacity": {
                    "type": "integer",
                    "minimum": 1
                },
                "className": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                }
            }
        },
        "responses.Response": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
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
	Title:            "Glofox Studio API",
	Description:      "API for managing studio classes and bookings",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
