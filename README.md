# ROBROKER - Mock Broker Application

Rest API mocking and intercepting in seconds. Replace the endpoint in the code and you are ready. It's that simple!

- [Installation and Setup](#installation-and-setup)
- [Create Broker](#create-broker)
    * [Request](#request)
    * [Response](#response)
        + [Created 201](#created-201)
- [Read Broker](#read-broker)
    * [Request](#request-1)
- [Register Broker](#register-broker)
    * [Request](#request)
    * [Response](#response)
        + [Created 201](#created-201)
- [Read Broker Catalog](#read-broker-catalog)
    * [Request](#request-1)
- [Update Broker](#update-broker)
    * [Patch Request](#patch-request)
    * [Response](#response-1)
        + [OK 200](#ok-200)
- [Delete Broker](#delete-broker)
    * [Request](#request-2)
    * [Response](#response-2)

## Installation and Setup

1. Clone the project
2. Import and Sync go.mod dependencies

The project contains a manifest.yaml file for cloud-foundry users.

1. Login to your cloud-foundry org and space.
2. Execute `cf push` from the repository directory.

## Create Broker

### Request

`POST` request to `https://robroker.<host>/broker`
The request body should contain the following JSON template:

Name and title your broker, then define a catalog according to the OSB spec.

https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#catalog-management

```json
{
  "name": "RobrokerName",
  "title": "RobrokerTitle",
  "provision": {
    "status": 201,
    "body": {}
  },
  "deprovision": {
    "status": 200,
    "body": {}
  },
  "bind": {
    "status": 201,
    "body": {}
  },
  "unbind": {
    "status": 200,
    "body": {}
  },
  "catalog": {
    "services": [
      {
        "name": "robroker-service",
        "description": "Provides an overview of any service instances and bindings that have been created by a platform.",
        "id": "891c24f4-5d66-423c-962e-ad73723d6a9b",
        "tags": [
          "robroker-broker"
        ],
        "bindable": true,
        "plan_updateable": true,
        "bindings_retrievable": true,
        "instances_retrievable": true,
        "metadata": {
          "shareable": true
        },
        "plans": [
          {
            "name": "small",
            "description": "A small instance of the service.",
            "free": true,
            "maintenance_info": {
              "version": "1.0.0"
            },
            "id": "da6274b3-debb-4518-bb95-8e90ca9e795b"
          },
          {
            "name": "large",
            "description": "A large instance of the service.",
            "free": true,
            "maintenance_info": {
              "version": "1.0.0"
            },
            "schemas": {
              "service_instance": {
                "create": {
                  "parameters": {
                    "$schema": "http://json-schema.org/draft-04/schema#",
                    "additionalProperties": false,
                    "type": "object",
                    "properties": {
                      "rainbow": {
                        "type": "boolean",
                        "default": false,
                        "description": "Follow the rainbow"
                      },
                      "name": {
                        "type": "string",
                        "minLength": 1,
                        "maxLength": 30,
                        "default": "This is a default string",
                        "description": "The name of the broker"
                      },
                      "color": {
                        "type": "string",
                        "enum": [
                          "red",
                          "amber",
                          "green"
                        ],
                        "default": "green",
                        "description": "Your favourite color"
                      },
                      "config": {
                        "type": "object",
                        "properties": {
                          "url": {
                            "type": "string"
                          },
                          "port": {
                            "type": "integer"
                          }
                        }
                      }
                    }
                  }
                },
                "update": {
                  "parameters": {
                    "$schema": "http://json-schema.org/draft-04/schema#",
                    "additionalProperties": false,
                    "type": "object",
                    "properties": {
                      "rainbow": {
                        "type": "boolean",
                        "default": false,
                        "description": "Follow the rainbow"
                      },
                      "name": {
                        "type": "string",
                        "minLength": 1,
                        "maxLength": 30,
                        "default": "This is a default string",
                        "description": "The name of the broker"
                      },
                      "color": {
                        "type": "string",
                        "enum": [
                          "red",
                          "amber",
                          "green"
                        ],
                        "default": "green",
                        "description": "Your favourite color"
                      },
                      "config": {
                        "type": "object",
                        "properties": {
                          "url": {
                            "type": "string"
                          },
                          "port": {
                            "type": "integer"
                          }
                        }
                      }
                    }
                  }
                }
              },
              "service_binding": {
                "create": {
                  "parameters": {
                    "$schema": "http://json-schema.org/draft-04/schema#",
                    "additionalProperties": false,
                    "type": "object",
                    "properties": {
                      "rainbow": {
                        "type": "boolean",
                        "default": false,
                        "description": "Follow the rainbow"
                      },
                      "name": {
                        "type": "string",
                        "minLength": 1,
                        "maxLength": 30,
                        "default": "This is a default string",
                        "description": "The name of the broker"
                      },
                      "color": {
                        "type": "string",
                        "enum": [
                          "red",
                          "amber",
                          "green"
                        ],
                        "default": "green",
                        "description": "Your favourite color"
                      },
                      "config": {
                        "type": "object",
                        "properties": {
                          "url": {
                            "type": "string"
                          },
                          "port": {
                            "type": "integer"
                          }
                        }
                      }
                    }
                  }
                }
              }
            },
            "id": "3665b124-e82f-4cdb-b5d8-236139c36d87"
          }
        ]
      }
    ]
  }
}
```

- `<method>` (`provision`/`deprovision`/`bind`/`unbind`) - (json object) contains the config (`status` and `body`) of
  the configured method.
- `status` - (int64) the expected response status of your `post`/`delete` requests
- `body` - the expected response body of your `post`/`delete` requests (can be a json format)

### Response

#### Created 201

```json
{
  "ID": "c2075f58-c5fb-45d4-9e54-a1ae9105e254",
  "Password": "admin",
  "URL": "https://robroker.<host>/broker/c2075f58-c5fb-45d4-9e54-a1ae9105e254/",
  "Username": "admin"
}
```

- `id` - (uuid) the unique id of your config, will be used for mocking your requests to the client and for setup/delete
- `<method>` (`post`/`delete`/`patch`) - (json object) contains the config (`status` and `body`) of the configured
  method.
- `status` - (int64) the expected response status of your `post`/`delete` requests
- `body` - the expected response body of your `post`/`delete` requests (can be a json format)

## Read Broker

### Request

`GET` request to `https://<host>/broker/<uuid>`
`uuid` - is the unique id of your mock config

**NOTE!**

You can get all configurations in your app memory using the `all=true` flag.

Example:
`GET` request to `https://<host>/broker?all=true`

## Update Broker

After creating a mock api config, you can change its settings.

This can be used when running a dynamic test that its configurations need to be changed on runtime.

### Patch Request

- `PATCH` request to `https://<host>/broker/<uuid>`
- `id` - (uuid) the unique id of your config, which received on the `Create Broker` step The request body should contain
  the following JSON template:

```json
{
  "catalog": {
    "services": [
      {
        "name": "service-manager",
        "plans": [
          {
            "name": "subaccount-admin"
          }
        ]
      }
    ]
  },
  "provision": {
    "status": 200,
    "body": {}
  },
  "deprovision": {
    "status": 200,
    "body": {}
  },
  "bind": {
    "status": 200,
    "body": {}
  },
  "unbind": {
    "status": 200,
    "body": {}
  }
}
```

- `<method>` (`provision`/`deprovision`/`bind`/`unbind`) - (json object) contains the config (`status` and `body`) of
  the configured method.
- `status` - (int64) the expected response status of your `post`/`delete` requests
- `body` - the expected response body of your `post`/`delete` requests (can be a json format)

### Response

#### OK 200

```json
{
  "id": "f137c451-e754-43ff-bb40-eb2a28b6db10",
  "catalog": {
    "services": [
      {
        "name": "service-manager111",
        "plans": [
          {
            "name": "subaccount-admin"
          }
        ]
      }
    ]
  },
  "name": "",
  "provision": {
    "status": 200,
    "body": {}
  },
  "deprovision": {
    "status": 200,
    "body": {}
  },
  "bind": {
    "status": 200,
    "body": {}
  },
  "unbind": {
    "status": 200,
    "body": {}
  }
}
```

- `id` - (uuid) the unique id of your config, will be used for mocking your requests to the client and for setup/delete
- `<method>` (`post`/`delete`/`patch`) - (json object) contains the config (`status` and `body`) of the configured
  method.
- `status` - (int64) the expected response status of your `post`/`delete` requests
- `body` - the expected response body of your `post`/`delete` requests (can be a json format)

## Delete Broker

When finishing with the mock api config, remember to delete it in order to avoid overloading your application in-memory.

### Request

`DELETE` request to `https://<host>/broker/<uuid>`
`uuid` - is the unique id of your mock config

**NOTE!**

You can delete all configurations in your app memory using the `all=true` flag.

Example:
`DELETE` request to `https://<host>/mockConfig?all=true`

### Response

`OK 200` on success

## Read Broker Catalog

### Request

`DELETE` request to `https://<host>/broker/<uuid>/v2/catalog`
`uuid` - is the unique id of your mock config

### Response

`OK 200` on success with the catalog provided earlier.
