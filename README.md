<h1 align="center"> Let's use Gechoplate 🚀</h1>

<p align="center">
    <img alt="Gechoplate" src="public/gechoplate.png" width="90%"/>
</p>

<div align="center">

<p>

  [![Go Doc](https://godoc.org/github.com//Akecel/gechoplate?status.svg)](https://godoc.org/github.com/Akecel/gechoplate)
  ![Build](https://github.com/Akecel/gechoplate/workflows/build/badge.svg?branch=master)
  [![Go Report](https://goreportcard.com/badge/github.com/Akecel/gechoplate)](https://goreportcard.com/report/github.com/Akecel/gechoplate)
  ![Version](https://img.shields.io/github/v/release/Akecel/gechoplate.svg)
  ![Licence](https://img.shields.io/badge/License-MIT-blue.svg)

</p>

</div>

## About Gechoplate

Gechoplate is a simple MVC boilerplate to design Rest APIs in Golang. It provides several basic features like a highly optimized HTTP router, a JWT authentication system and a full-featured QueryBuilder.

>This boilerplate has been designed to be used by everyone and especially for developers new to the Go language, feel free to use it in your personnal or school projects !

### Features

* Highly optimized HTTP router & built-in Middleware with [Echo Framework](https://github.com/labstack/echo)
* Easy and clean configuration file system with [Viper](https://github.com/spf13/viper)
* JWT Authentification system with [jwt-go](https://github.com/dgrijalva/jwt-go)
* MySQL/PostgreSQL support including QueryBuilder, migration and seeding with [GORM](https://github.com/go-gorm/gorm)
* Easy and complete request data validation using [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* All in one [Docker compose](https://docs.docker.com/compose/install/) with a Go build and a MySQL server
* Friendly automation tool for project management with *Make*

## Requirement

Gechoplate comes with an all-in-one docker-compose including a Go build, a MySQL or PostgreSQL server.

To use the boilerplate in an optimal way, docker is required.

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/install/)

## Use this boilerplate

With [GitHub CLI](https://cli.github.com/) :

```
$ gh repo create [name] --template Akecel/gechoplate
```

Or use the button **[Use this template](https://github.com/Akecel/gechoplate/generate)** above

You can also clone this repository :

```
$ git clone https://github.com/Akecel/gechoplate.git
```

## Configuration

Gechoplate is using [Viper](https://github.com/spf13/viper) to provide a complete configuration file system.

To use the application, you will need to generate an environment file using :

```
$ make env
```
*If an environment file already exists, it will be replaced by a new one, **be careful not to overwrite your configuration.***

To be beginners-friendly Docker-compose and Gechoplate use the same environment variables.

This way, you only have the *.env* file to configure :

```
APP_URL=http:localhost:1323
APP_NAME=Gechoplate
DB_HOST=db
DB_NAME=database
DB_USER=user
DB_PASSWORD=password
DB_PORT=3306
```

By default Gechoplate uses MySQL, however you can use PostgreSQL very easily if you wish to do so : 

* Comment/uncomment the different SQL containers in the *docker-compose* file depending on the database system you want to use :

```yml
  # MySQL Support
  db:
    image: mysql:8.0
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: ${DB_NAME}
      MYSQL_USER: ${DB_USER}
      MYSQL_PASSWORD: ${DB_PASSWORD}
    volumes:
      - db-data:/var/lib/mysql
    ports:
      - 3306:3306

  # PostgreSQL Support
  # db:
  #   image: postgres:12.2-alpine
  #   environment:
  #     POSTGRES_DB: ${DB_NAME}
  #     POSTGRES_USER: ${DB_USER}
  #     POSTGRES_PASSWORD: ${DB_PASSWORD}
  #     PGDATA: /data/postgres
  #   volumes:
  #     - db-data:/data/postgres
  #   ports:
  #     - 5432:5432
```
*Don't forgret to change the "DB_PORT" in your .env file if you use PostgreSQL*

* You will also have to modify the Gorm connector : 

```go
// MySQL connexion
Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// PostgreSQL connexion
//Gorm, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

Finally, to fully use this boilerplate for your projects, remember to change the name of the module in the go.mod file and to modify this name in all the project imports.

```go
module your_project
```

```go
package controllers

import (

  db "your_project/database"
  "your_project/helpers"
  "your_project/models"
  
)
```

## Usage

### Makefile

Gechoplate use a *Makefile* to manage all commands of the project.

You can display the list of commands with :

```
$ make help
Usage: make <command>

Commands:
  help                   Provides help information on available commands.
  compose/build          Build all Docker images of the project.
  compose/up             Start all containers (in the background).
  compose/down           Stops and deletes containers.
  compose/purge          Stops and deletes containers, volumes, images and networks.
  compose/rebuild        Rebuild all Docker images of the project.
  urls                   Get projects URL.
  env                    Generate env file.
```

### Run application

Start the project with :

```
$ make compose/up
Creating network "gechoplate_default" with the default driver
Creating gechoplate_db_1 ... done
Creating gechoplate_go_1 ... done


 _____           _                 _       _
|  __ \         | |               | |     | |
| |  \/ ___  ___| |__   ___  _ __ | | __ _| |_ ___
| | __ / _ \/ __|  _ \ / _ \|  _ \| |/ _  | __/ _ \
| |_\ \  __/ (__| | | | (_) | |_) | | (_| | ||  __/
 \____/\___|\___|_| |_|\___/|  __/|_|\__ _|\__\___|
                            | |
                            |_|

```

You can now access the API: [http://localhost:1323/](http://localhost:1323/).

## Test

Gechoplate has a *tests* directory containing all the test files for each controller.\
The test files are generated using [gotests](https://github.com/cweill/gotests).

Simply use the test go command to run your tests :

```
$ cd tests
$ go test
PASS
ok      gechoplate/tests        0.029s
```

## Contributing

Please read [CONTRIBUTING.md](https://github.com/Akecel/gechoplate/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* [**Akecel**](https://github.com/Akecel) - *Template author*
* [**Hugo-T**](https://github.com/T-Hugo) - *Development environment author*

See also the list of [contributors](https://github.com/Akecel/gechoplate/graphs/contributors) who participated in this project.

## Licence

This project is licensed under the [MIT License](https://opensource.org/licenses)  - see the [LICENSE.md](https://github.com/Akecel/gechoplate/blob/master/LICENSE) file for details.

## Show your support

Give a ⭐️ if this project helped you!
