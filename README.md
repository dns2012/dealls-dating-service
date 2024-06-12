# Dealls Dating Service

RESTful API prototye for Dating Service required by Dealls.

## Table of Contents

- [About](#about)
- [Features](#features)
- [Depedency](#depedency)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## About

This repository is a RESTful API prototype of a dating application called Dealls Dating Service. The purpose of creating this prototype is for the technical test required by Dealls.

## Features

List of key features or functionalities of the project:

- Authentication (Sign In / Sign Up)
- List potential users for date
- Set user preference (Pass / Like)
- Show my profile
- List of packages
- Purchase (Order) package for more swap freedom

## Depedency
Since this application uses Docker, you can run it simply with Docker. Here is the list of dependencies:
- Docker* (Required in the machine host)
- GoLang 1.21.3 (Optional in the machine host)
- MySQL 8 (Optional in the machine host)
  Protobuf / Buf (Optional in the machine host)
- CMAKE* or MAKE* (Required in the machine host)


## Installation

First of all setup the environment, copy file `.env.example` to `.env`, by default the value of the example is ready
to use in real case.

Next, start docker container for the service

```bash
$ make start
```

if there is no CMAKE / MAKE installed, do it by
```bash
$ docker-compose up -d
```

After the service is started, try to migrate default database structure and seeder
```bash
$ make migrate
```
if there is no CMAKE / MAKE installed, do it by
```bash
$ docker-compose exec -it httpd go run database/migration.go
```

## Usage

After installation and setup is success, this services will open 2 ports in the machine :
- 8080 (gRpc)
- 8081 (gRpc Gateway)
since the requirement is RESTful API, the common usage must be in `http://localhost:8081` as the entrypoint
  

## Documentation

By today there is still no automatic RESTful API documentation serve in browser, but there is generated Open API Swagger documentation
inside `docs/schema.swagger.json` which easily copy and paste to the `https://editor.swagger.io` or import through Postman. 
There is also file `docs/schema.har` which possible to imported as documentation collection to the Postman / Insomnia