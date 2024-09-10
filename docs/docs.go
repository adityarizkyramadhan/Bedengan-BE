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
            "name": "Aditya Rizky Ramadhan",
            "email": "adityarizky1020@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/perlengkapan": {
            "get": {
                "description": "Menampilkan semua data perlengkapan",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Perlengkapan"
                ],
                "summary": "Menampilkan semua data perlengkapan",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            },
            "post": {
                "description": "Membuat data perlengkapan baru",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Perlengkapan"
                ],
                "summary": "Membuat data perlengkapan baru",
                "parameters": [
                    {
                        "description": "Perlengkapan data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PerlengkapanInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/perlengkapan/{id}": {
            "get": {
                "description": "Menampilkan data perlengkapan berdasarkan id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Perlengkapan"
                ],
                "summary": "Menampilkan data perlengkapan berdasarkan id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID Perlengkapan",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            },
            "put": {
                "description": "Memperbarui data perlengkapan berdasarkan id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Perlengkapan"
                ],
                "summary": "Memperbarui data perlengkapan berdasarkan id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID Perlengkapan",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Perlengkapan data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PerlengkapanInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            },
            "delete": {
                "description": "Menghapus data perlengkapan berdasarkan id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Perlengkapan"
                ],
                "summary": "Menghapus data perlengkapan berdasarkan id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID Perlengkapan",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Find user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Find user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            },
            "put": {
                "description": "Update user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "get": {
                "description": "Logout user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Logout user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/resend/{email}": {
            "get": {
                "description": "Resend email OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Resend email OTP",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/verify/{otp}": {
            "get": {
                "description": "Verify OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Verify OTP",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OTP",
                        "name": "otp",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.SuccessResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.PerlengkapanInput": {
            "type": "object",
            "required": [
                "deskripsi",
                "harga",
                "nama",
                "stok"
            ],
            "properties": {
                "deskripsi": {
                    "type": "string"
                },
                "harga": {
                    "type": "integer"
                },
                "nama": {
                    "type": "string"
                },
                "stok": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "city": {
                    "description": "Kota tempat tinggal user",
                    "type": "string"
                },
                "email": {
                    "description": "Email bersifat unik dan tidak boleh kosong",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "description": "Name tidak boleh kosong",
                    "type": "string"
                },
                "province": {
                    "description": "Provinsi tempat tinggal user",
                    "type": "string"
                },
                "role": {
                    "description": "Role merupakan enum yang berisi \"admin\" dan \"user\" dan not null",
                    "type": "string"
                }
            }
        },
        "model.UserCreate": {
            "type": "object",
            "required": [
                "city",
                "confirm_password",
                "email",
                "name",
                "password",
                "province"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                }
            }
        },
        "model.UserLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UserUpdate": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                }
            }
        },
        "utils.ErrorResponseData": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Example: nil"
                },
                "message": {
                    "description": "The error message.\nExample: \"Invalid request\"",
                    "type": "string"
                }
            }
        },
        "utils.SuccessResponseData": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "The actual data returned."
                },
                "message": {
                    "description": "The success message.\nExample: \"success\"",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.perkemahanbedengan.com",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Islamind API",
	Description:      "This is API documentation for Islamind. You can use the API with the following hosts:\n- Production: `api.perkemahanbedengan.com`",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
