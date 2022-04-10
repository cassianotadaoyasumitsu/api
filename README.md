# API

GoLang

MySQL

Validations

Middlewares

JSON Web Token / Token Validation

bcrypt

CRUD

This API was build for educational propose only! Please feel free to clone and improve it.

Allow us to create an account, publish posts, follow users, give likes and so far so on. 


Example of:

````

// 20220410233629
// https://raw.githubusercontent.com/cassianotadaoyasumitsu/api/master/Devbook.postman_collection.json

{
  "info": {
    "_postman_id": "2b8d8b1c-43aa-4968-90e3-6aa1e199667b",
    "name": "Devbook",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "find users",
      "protocolProfileBehavior": {
        "disableBodyPruning": true
      },
      "request": {
        "method": "GET",
        "header": [
          
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"email\": \"user1@email.com\",\n    \"password\": \"123456\"\n}",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "localhost:5001/users",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "users"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "login",
      "request": {
        "method": "POST",
        "header": [
          
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"email\": \"cassiano@email.com\",\n    \"password\": \"123456\"\n}",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "localhost:5001/login",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "login"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "add users",
      "request": {
        "method": "POST",
        "header": [
          
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"name\": \"Cassiano\",\n    \"nick\": \"Cassi\",\n    \"email\": \"cassiano@email.com\",\n    \"password\": \"123456\"\n}",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "localhost:5001/users",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "users"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "update",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ3MTE4OTc4LCJ1c2VySUQiOjh9.I0PEqwZEK4cCIsUcCN6dTSYdJ-cgV1V0_9BFLkM1tfA",
              "type": "string"
            }
          ]
        },
        "method": "PUT",
        "header": [
          
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"name\": \"Mieko\",\n    \"nick\": \"olivia\",\n    \"email\": \"olivia@email.com\",\n    \"password\": \"123456\"\n}"
        },
        "url": {
          "raw": "localhost:5001/users/1",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "users",
            "1"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "delete",
      "request": {
        "method": "DELETE",
        "header": [
          
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"name\": \"Mieko\",\n    \"nick\": \"olivia\",\n    \"email\": \"olivia@email.com\",\n    \"password\": \"123456\"\n}"
        },
        "url": {
          "raw": "localhost:5001/users/1",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "users",
            "1"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "follow",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ3Mjg0MTQwLCJ1c2VySUQiOjR9.4KebhGW84MlVvvIOUHLBSxuntzOu_X24kxM7XRgnGxI",
              "type": "string"
            }
          ]
        },
        "method": "POST",
        "header": [
          
        ],
        "url": {
          "raw": "localhost:5001/users/1/follow",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "users",
            "1",
            "follow"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "unfollow",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ3Mjg0MTQwLCJ1c2VySUQiOjR9.4KebhGW84MlVvvIOUHLBSxuntzOu_X24kxM7XRgnGxI",
              "type": "string"
            }
          ]
        },
        "method": "POST",
        "header": [
          
        ],
        "body": {
          "mode": "raw",
          "raw": "{}"
        },
        "url": {
          "raw": "localhost:5001/users/1/unfollow",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "users",
            "1",
            "unfollow"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "followers",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4MDU5MDE3LCJ1c2VySUQiOjR9.yCDow1LZZcezHx_OD8ibwOxGklddc8rcqzd-UXqlt6s",
              "type": "string"
            }
          ]
        },
        "method": "GET",
        "header": [
          
        ],
        "url": {
          "raw": "localhost:5001/users/2/followers",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "users",
            "2",
            "followers"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "following",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4MDU5MDE3LCJ1c2VySUQiOjR9.yCDow1LZZcezHx_OD8ibwOxGklddc8rcqzd-UXqlt6s",
              "type": "string"
            }
          ]
        },
        "method": "GET",
        "header": [
          
        ],
        "url": {
          "raw": "localhost:5001/users/4/following",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "users",
            "4",
            "following"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "password-update",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4MDYyOTY0LCJ1c2VySUQiOjR9._MYxvJaJ7X4KaZQjWUlsgJ7cM5z3F0hoqJrCXEXnCPs",
              "type": "string"
            }
          ]
        },
        "method": "POST",
        "header": [
          
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"new\": \"654321\",\n    \"old\": \"123456\"\n}"
        },
        "url": {
          "raw": "localhost:5001/users/4/password-update",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "users",
            "4",
            "password-update"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "create posts",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4MTMzMjYyLCJ1c2VySUQiOjR9.wrrCDHoWIyk6c-yEGBnq_NUNwWuXDO8lREut8a-QATI",
              "type": "string"
            }
          ]
        },
        "method": "POST",
        "header": [
          
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"title\": \"My first post\",\n    \"content\": \"Hello World Cassiano\"\n}"
        },
        "url": {
          "raw": "localhost:5001/posts",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "posts"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "search post",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4MTMzMjYyLCJ1c2VySUQiOjR9.wrrCDHoWIyk6c-yEGBnq_NUNwWuXDO8lREut8a-QATI",
              "type": "string"
            }
          ]
        },
        "method": "GET",
        "header": [
          
        ],
        "url": {
          "raw": "localhost:5001/posts/2",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "posts",
            "2"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "search posts",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4MTUwMjA2LCJ1c2VySUQiOjF9.7nWRwTSNc3JZB7La2NnX0Q8K5DZKF-XgjWubR9CQawA",
              "type": "string"
            }
          ]
        },
        "method": "GET",
        "header": [
          
        ],
        "url": {
          "raw": "localhost:5001/posts",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "posts"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "update post",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4NTA0OTM5LCJ1c2VySUQiOjR9.acpRP5-npA1D7IH9jKVK-LbQp6iA13NkHqTa_EdbhZ0",
              "type": "string"
            }
          ]
        },
        "method": "PUT",
        "header": [
          
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"title\": \"New Title\",\n    \"content\": \"New content\"\n}"
        },
        "url": {
          "raw": "localhost:5001/posts/4",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "posts",
            "4"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "delete post",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4NTA0OTM5LCJ1c2VySUQiOjR9.acpRP5-npA1D7IH9jKVK-LbQp6iA13NkHqTa_EdbhZ0",
              "type": "string"
            }
          ]
        },
        "method": "DELETE",
        "header": [
          
        ],
        "url": {
          "raw": "localhost:5001/posts/1",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "posts",
            "1"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "search posts by user",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4NTA2OTk4LCJ1c2VySUQiOjN9.LEYH3F1g916aLomhje_z_at75OouXId48iKAceO5kHQ",
              "type": "string"
            }
          ]
        },
        "method": "GET",
        "header": [
          
        ],
        "url": {
          "raw": "localhost:5001/users/4/posts",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "users",
            "4",
            "posts"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "likes",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4NTg3NzIwLCJ1c2VySUQiOjR9.vgfXUfDgzN8kXjuwlZrpuB6LnldOR_vvzASf7PExST8",
              "type": "string"
            }
          ]
        },
        "method": "POST",
        "header": [
          
        ],
        "url": {
          "raw": "localhost:5001/posts/10/likes",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "posts",
            "10",
            "likes"
          ]
        }
      },
      "response": [
        
      ]
    },
    {
      "name": "unlikes",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVzIjoxNjQ4NTg3NzIwLCJ1c2VySUQiOjR9.vgfXUfDgzN8kXjuwlZrpuB6LnldOR_vvzASf7PExST8",
              "type": "string"
            }
          ]
        },
        "method": "POST",
        "header": [
          
        ],
        "url": {
          "raw": "localhost:5001/posts/10/unlikes",
          "host": [
            "localhost"
          ],
          "port": "5001",
          "path": [
            "posts",
            "10",
            "unlikes"
          ]
        }
      },
      "response": [
        
      ]
    }
  ]
}

