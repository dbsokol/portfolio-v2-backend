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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/contacts": {
            "get": {
                "description": "List all contacts in the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contacts"
                ],
                "summary": "List all contacts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "enum": [
                                "CREATED_AT",
                                "-CREATED_AT",
                                "NAME",
                                "-NAME"
                            ],
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Ordering",
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name insensitive contains",
                        "name": "nameIContains",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/contact.Contact"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new contact to the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contacts"
                ],
                "summary": "Create a new contact",
                "parameters": [
                    {
                        "description": "Contact data",
                        "name": "contact",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contact.CreateContactRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/contact.Contact"
                        }
                    }
                }
            }
        },
        "/educations": {
            "get": {
                "description": "List all education records from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "educations"
                ],
                "summary": "List all education records",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "enum": [
                                "BACHELORS",
                                "MASTERS",
                                "DOCTORATE"
                            ],
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Degrees",
                        "name": "degrees",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "enum": [
                                "START_DATE",
                                "-START_DATE",
                                "END_DATE",
                                "-END_DATE",
                                "CREATED_AT",
                                "-CREATED_AT",
                                "DEGREE",
                                "-DEGREE"
                            ],
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Ordering",
                        "name": "ordering",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/education.Education"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new education record to the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "educations"
                ],
                "summary": "Create a new education record",
                "parameters": [
                    {
                        "description": "Education data",
                        "name": "education",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/education.CreateEducationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/education.Education"
                        }
                    }
                }
            }
        },
        "/educations/{uuid}/": {
            "delete": {
                "description": "Soft delete an education record from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "educations"
                ],
                "summary": "Delete an education record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Education UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Partially update an education record in the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "educations"
                ],
                "summary": "Partially update an education record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Education UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Education data",
                        "name": "education",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/education.PartialUpdateEducationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/education.Education"
                        }
                    }
                }
            }
        },
        "/skills": {
            "get": {
                "description": "Retrieve a list of skills from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "skills"
                ],
                "summary": "List skills",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit the number of results returned",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset the results returned",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by start date greater than or equal to (YYYY-MM-DD)",
                        "name": "startDateGTE",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by start date less than or equal to (YYYY-MM-DD)",
                        "name": "startDateLTE",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "enum": [
                                "LANGUAGE",
                                "FRAMEWORK",
                                "CLOUD"
                            ],
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Filter by skill types (multi-select)",
                        "name": "types",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "enum": [
                                "CREATED_AT",
                                "-CREATED_AT",
                                "NAME",
                                "-NAME",
                                "START_DATE",
                                "-START_DATE"
                            ],
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Order by fields (multi-select)",
                        "name": "ordering",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/skill.Skill"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new skill to the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "skills"
                ],
                "summary": "Create a new skill",
                "parameters": [
                    {
                        "description": "Skill data",
                        "name": "skill",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/skill.CreateSkillRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/skill.Skill"
                        }
                    }
                }
            }
        },
        "/skills/{uuid}/": {
            "delete": {
                "description": "Delete a skill from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "skills"
                ],
                "summary": "Delete a skill",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Skill UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Partially update a skill in the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "skills"
                ],
                "summary": "Partially update a skill",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Skill UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Skill data",
                        "name": "skill",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/skill.PartialUpdateSkillRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/skill.Skill"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "contact.Contact": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "contact.CreateContactRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "url": {
                    "description": "Optional",
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "education.CreateEducationRequest": {
            "type": "object",
            "required": [
                "degree",
                "endDate",
                "institution",
                "major",
                "startDate"
            ],
            "properties": {
                "degree": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "institution": {
                    "type": "string"
                },
                "major": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                }
            }
        },
        "education.Education": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "degree": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "institution": {
                    "type": "string"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "major": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "education.PartialUpdateEducationRequest": {
            "type": "object",
            "properties": {
                "degree": {
                    "description": "Optional",
                    "type": "string"
                },
                "endDate": {
                    "description": "Optional",
                    "type": "string"
                },
                "institution": {
                    "description": "Optional",
                    "type": "string"
                },
                "major": {
                    "description": "Optional",
                    "type": "string"
                },
                "startDate": {
                    "description": "Optional",
                    "type": "string"
                }
            }
        },
        "skill.CreateSkillRequest": {
            "type": "object",
            "required": [
                "name",
                "startDate",
                "type"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "skill.PartialUpdateSkillRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Optional",
                    "type": "string"
                },
                "startDate": {
                    "description": "Optional",
                    "type": "string"
                },
                "type": {
                    "description": "Optional",
                    "type": "string"
                }
            }
        },
        "skill.Skill": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "uuid": {
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
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Portfolio API",
	Description:      "API documentation for the Portfolio project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}