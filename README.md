<h1 align="center">AyoBelajar LMS App - Backend</h1>

This API is created by me for build my personal project that will be helpful in future. This backend api using Golang with Gin and GORM package. [More about Golang](https://go.dev/)

## Built With

[![Go/Golang](https://img.shields.io/badge/Golang-1.20-cyan.svg?style=rounded-square)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-v.1.10-blue.svg?style=rounded-square)](https://gin-gonic.com/docs/)
[![GORM](https://img.shields.io/badge/Gorm-v.1.25-blue.svg?style=rounded-square)](https://gorm.io/)


## Requirements

1. <a href="https://go.dev/">Go</a>
2. <a href="https://gin-gonic.com/docs/">Gin</a>
3. <a href="https://gorm.io/">GORM</a>
4. <a href="https://www.mysql.com/">MySQL</a>

## How to run the app ?

1. Clone this project
2. Open app's directory in CMD or Terminal
3. Turn on Web Server and MySQL can using Third-party tool like xampp, etc.
4. Create a database with the name #nama_database, and Import file sql to **phpmyadmin**
5. Open Postman desktop application or Chrome web app extension that has installed before
6. Choose HTTP Method and enter request url.(ex. localhost:3600/)
7. You can see all the end point [here](https://documenter.getpostman.com/view/14780095/2sAXjM3XAX)
8. Type `go run main.go` to activated the server.

## Set up project

1. Create a file ".env" then setup environment variable like code below:

```
APP_PORT=<GIN LISTENING PORT>

DATABASE_HOSTNAME=<YOUR_DB_HOSTNAME>
DATABASE_PORT=<YOUR_DB_PORT>
DATABASE_USERNAME=<YOUR_DB_USERNAME>
DATABASE_PASSWORD=<YOUR_DB_PASSWORD>
DATABASE_NAME=<YOUR_DATABASE_NAME>

JWT_PRIVATE_KEY=<YOUR_SECRET_KEY_JWT>
```

## Feature

1. Login and Register User
2. CRUD For Users (For Admin And User)
3. CRUD For Course (For Admin)
4. CRUD For Course Category (For Admin)
5. CRUD For Course Cart(For Admin & User)
6. CRUD For Course Invoice Payment(For Admin & User)


## License

© [Muhammad Akbar Saladin Siregar](https://github.com/akbarsaladin36/)
