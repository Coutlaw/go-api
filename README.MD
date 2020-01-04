# go-api

My first experiamental golang API that allows users to store contacts. Users are authenticated with a JWT on their requests. I used `gorilla/mux` for routing, and Postgresql running in Google Cloud Platform for a db. There is a `.env` file for managing connection details. The project was rebuilt by me from this [Medium Article](https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b) to teach myself the basics of a golang API

# Organization
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
├── LICENSE.md
├── README.MD
| 
├── .env (not included in repo)
```
`app` contains our auth functions and error handling.
`controllers` provides all the controller layer operations. `models` provide model shapes, content/request validation and operational logic with the DB. `utils` provides json serialization for responses and message handling functions.

# Enviornment
Not included is a `.env` file that will need to be formatted like so :
```
db_name = gocontacts
db_pass = pAsSwOrD
db_user = postgres-user
db_type = postgres
db_host = localhost
db_port = 5432
token_password = jwtPassword
database_url = postgres-url
```
The content in this file will provide all DB connection details for the application to run.

# Running

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
	


# Routes
Creating a user:
`api/user/new`

Notes: The user will be provided a JWT in return to use for future requested based on the info they provided, the user's email must be unique in the DB and passwords > 6 characters.

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

# Logging
More content to come

# Portability
App will be dockerised in the near future, more content to follow

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
