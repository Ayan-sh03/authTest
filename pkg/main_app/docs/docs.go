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
        "/user/login": {
            "post": {
                "description": "Allows users to login into their account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login route",
                "parameters": [
                    {
                        "description": "User's email and password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/doc_model.Login"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON data, Invalid Email",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Please Verify Your Account, Invalid Credentials",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User is not registered",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Please Verify Your Account",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/otp": {
            "post": {
                "description": "Allows users to validate OTP and complete the registration process.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Validation route",
                "parameters": [
                    {
                        "description": "User's email address and otp",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/doc_model.OTP"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/doc_model.OTP_successResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON data, Invalid Email",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid OTP, User Already Verified",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Allows users to create a new account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register route",
                "parameters": [
                    {
                        "description": "User's firstname, lastname, middlename, email, password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/doc_model.Register"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON data, Invalid Email",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Please provide with sufficient credentials",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error, Error in inserting the document, Error in hashing password, Error While generating OTP",
                        "schema": {
                            "$ref": "#/definitions/doc_model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "doc_model.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "doc_model.Login": {
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
        "doc_model.OTP": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "otp": {
                    "type": "string"
                }
            }
        },
        "doc_model.OTP_successResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "doc_model.Register": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "middlename": {
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
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Registration API",
	Description:      "This is a registration api for an application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
