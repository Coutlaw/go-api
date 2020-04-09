# **ROLODEX, an API in Go**

My first experimental gol API that allows users to store contacts like a rolodex. Users are authenticated with a JWT on their requests. I used `gorilla/mux` for routing, and Postgresql running in Google Cloud Platform for a db. There is a `.env` file for managing connection details. The project was built to teach myself the basics of a Go API

# **ORGANIZATION**
The project hierarchy:
```
go-api
├── app
│   ├── auth.go
│   └── errors.go
├── controllers
│   ├── authControllers.go
│   └── contactsControllers.go
├── models
│   ├── accounts.go
│   ├── base.go
│   └── contact.go
├── utils
|    └── util.go
├── go.mod
├── go.sum
├── main.go
├── Dockerfile
├── docker-compose.yml
├── LICENSE.md
├── README.MD
| 
├── .env (not included in repo)
```
`app` contains our auth functions and error handling.
`controllers` provides all the controller layer operations. `models` provide model shapes, content/request validation and operational logic with the DB. `utils` provides json serialization for responses and message handling functions.

# **ENVIORNMENT**
Not included is a `.env` file that will need to be formatted like so :
```
# info to point at a remote sql server
db_name = gocontacts
db_pass = pAsSwOrD
db_user = postgres-user
db_type = postgres
db_host = localhost
db_port = 5432
token_password = jwtPassword
database_url = postgres-url

# Used by pgadmin service 
PGADMIN_DEFAULT_EMAIL=live@admin.com
PGADMIN_DEFAULT_PASSWORD=password

# info to point at the sql running in docker
TEST_DB_HOST=fullstack-postgres
# TEST_DB_HOST=127.0.0.1                        # when running the app without docker 
TEST_DB_DRIVER=postgres
TEST_API_SECRET=someSecretForJWT                # can be whatever your want
TEST_DB_USER=cass
TEST_DB_PASSWORD=password
TEST_DB_NAME=fullstack_api_test
TEST_DB_PORT=5432
```
The content in this file will provide all DB connection details for the application to run. You will need to update the properties in base.go to tell it which set of variables you want to use for postgresql, by default I have pointed it at the sql that will be running in docker.

# **RUNNING IN DOCKER**
Build the docker container from compose file and run the application with its own Postgresql image
```
docker-compose up
```
This will start the application, and you should see that it created all the required tables for the app in postgresql. Sometimes it can fail to connect if it is your first build, simply stop the container and re-run the compose up command. you can add `-d` flag to the end of the up command if you want it to silence all the logs.

To stop the container
```
docker-compose down
```

To destroy the images when you stop the container
```
docker-compose down --rmi all
```

To remove the Postgresql volume and wipe the DB
```
docker volume ls          # list all the volumes
docker volume rm {project-name}_database_postgre
```

# **RUNNING WITHOUT DOCKER**

Note: dependencies require go version `go1.12.0` or greater
Required dependencies to install with `go get {dependency}
```
github.com/lib/pq
github.com/gorilla/mux 
github.com/dgrijalva/jwt-go                          
github.com/joho/godotenv 
```
Steps build and run
1. Build the project

	`go build *.go`

2. Run the main file

	`go run main.go`

The app will run on port 8000 and you should see a line that looks like this: 
```
host={.env/db_host} user={.env/db_user} dbname={.env/db_name} sslmode=disable password={.env/db_pass}

Listening on port: 8000 
```
	


# **ROUTES**

## `api/user/new`
------
### POST
The user will be provided a JWT in return to use for future requested based on the info they provided, the user's email must be unique in the DB and passwords > 6 characters.

Request:
```
curl --request POST \
  --url http://localhost:8000/api/user/new \
  --header 'content-type: application/json' \
  --data '{
	"email" : "test@testytest.com",
	"password" : "PaSsWoRd"
}'
```
Response:
```
// Success
{
  "account": {
    "ID": 2,
    "CreatedAt": "2020-01-04T11:38:58.149453-05:00",
    "UpdatedAt": "2020-01-04T11:38:58.149453-05:00",
    "DeletedAt": null,
    "email": "test1@testytest.com",
    "password": "",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjJ9.bHzCP2yt40ISTCT3bgDju4T-qQfKz1LfbOgGB1qCS68"
  },
  "message": "Account has been created",
  "status": true
}

// Error
{
  "message": "Email address already in use by another user.",
  "status": false
}
```
---
## `api/user/login`
---
### POST

Notes: The user will be provided a JWT in return to use for future requested based on the info they provided, the user's email and password must match what they used when they created their account

Request:
```
curl --request POST \
  --url http://localhost:8000/api/users/login \
  --header 'content-type: application/json' \
  --data '{
	"email" : "test@testytest.com",
	"password" : "PaSsWoRd"
}'
```
Response:
```
// Success
{
  "account": {
    "ID": 1,
    "CreatedAt": "2020-01-04T15:56:29.252043Z",
    "UpdatedAt": "2020-01-04T15:56:29.252043Z",
    "DeletedAt": null,
    "email": "test@testytest.com",
    "password": "",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjF9.qq_02tATMUBkVUjI0cMplK8iN9Lv0ce_kZGtE3BuOtw"
  },
  "message": "Logged In",
  "success": true
}

// Error
{
  "message": "Invalid login credentials. Please try again",
  "success": false
}
```
---

## `api/user/contacts`
---

### POST
Creates a new contact for a user. The user will have to attache an authentication header with a bearer token from their JWT.

Request:
```
curl --request POST \
  --url http://localhost:8000/api/users/contacts \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjF9.qq_02tATMUBkVUjI0cMplK8iN9Lv0ce_kZGtE3BuOtw' \
  --header 'content-type: application/json' \
  --data '{
	"name" : "contact num4boi",
	"phone" : "8033332227"
}'
```
Response:
```
// Success
{
  "contact": {
    "ID": 5,
    "CreatedAt": "2020-01-06T23:43:47.655196-05:00",
    "UpdatedAt": "2020-01-06T23:43:47.655196-05:00",
    "DeletedAt": null,
    "name": "contact num4boi",
    "phone": "8033332227",
    "user_id": 1
  },
  "message": "success",
  "success": true
}

// Error
{
  "message": "Phone number should be on the payload",
  "success": false
}
```

### GET
Gets all contacts for a user. The user will have to attache an authentication header with a bearer token from their JWT.

Request:
```
curl --request GET \
  --url http://localhost:8000/api/users/contacts \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjF9.qq_02tATMUBkVUjI0cMplK8iN9Lv0ce_kZGtE3BuOtw'
```
Response:
```
// Success
{
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2020-01-07T02:14:51.863213Z",
      "UpdatedAt": "2020-01-07T02:14:51.863213Z",
      "DeletedAt": null,
      "name": "contact name",
      "phone": "8033332222",
      "user_id": 1
    },
    {
      "ID": 2,
      "CreatedAt": "2020-01-07T02:47:21.328672Z",
      "UpdatedAt": "2020-01-07T02:47:21.328672Z",
      "DeletedAt": null,
      "name": "contact name",
      "phone": "8033332222",
      "user_id": 1
    },
    {
      "ID": 3,
      "CreatedAt": "2020-01-07T03:37:17.741761Z",
      "UpdatedAt": "2020-01-07T03:37:17.741761Z",
      "DeletedAt": null,
      "name": "contact name",
      "phone": "8033332222",
      "user_id": 1
    }
  ],
  "message": "success",
  "success": true
}

// Error
{
  "message": "user not valid",
  "success": false
}
```

### DELETE
DELETES a single contact for a user by using the contact ID. The user will have to attache an authentication header with a bearer token from their JWT.

Request:
```
curl --request DELETE \
  --url http://localhost:8000/api/users/contacts/2 \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjF9.qq_02tATMUBkVUjI0cMplK8iN9Lv0ce_kZGtE3BuOtw'
```
Response:
```
// Success
{
  "message": "success",
  "success": true
}

// Error
{
  "message": "user not valid",
  "success": false
}
```

---

## `api/user/contacts/{ID}`
---

### GET
Gets a single contact for a user by using the contact ID. The user will have to attache an authentication header with a bearer token from their JWT.

Request:
```
curl --request GET \
  --url http://localhost:8000/api/users/contacts/2 \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjF9.qq_02tATMUBkVUjI0cMplK8iN9Lv0ce_kZGtE3BuOtw'
```
Response:
```
// Success
{
  "data": {
    "ID": 2,
    "CreatedAt": "2020-01-07T02:47:21.328672Z",
    "UpdatedAt": "2020-01-07T02:47:21.328672Z",
    "DeletedAt": null,
    "name": "contact name",
    "phone": "8033332222",
    "user_id": 1
  },
  "message": "success",
  "success": true
}

// Error
{
  "message": "user not valid",
  "success": false
}
```

### DELETE
DELETES a single contact for a user by using the contact ID. The user will have to attache an authentication header with a bearer token from their JWT.

Request:
```
curl --request DELETE \
  --url http://localhost:8000/api/users/contacts/2 \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjF9.qq_02tATMUBkVUjI0cMplK8iN9Lv0ce_kZGtE3BuOtw'
```
Response:
```
// Success
{
  "message": "success",
  "success": true
}

// Error
{
  "message": "user not valid",
  "success": false
}
```

# Logging
More content to come as I add better system logging


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
