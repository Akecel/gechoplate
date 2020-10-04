<h1 align="center"> Let's use Gechoplate 🚀</h1>

<p align="center">
    <img alt="Gechoplate" src="public/gechoplate.png" width="90%"/>
</p>

<div align="center">

<p>

  [![Go Doc](https://godoc.org/github.com//Akecel/gechoplate?status.svg)](https://godoc.org/github.com/Akecel/gechoplate)
  ![Build](https://github.com/Akecel/gechoplate/workflows/build/badge.svg?branch=master)
  [![Go Report](https://goreportcard.com/badge/github.com/Akecel/gechoplate)](https://goreportcard.com/report/github.com/Akecel/gechoplate)
  ![Version](https://img.shields.io/badge/version-0.1-blue.svg?cacheSeconds=2592000)
  ![Licence](https://img.shields.io/badge/License-MIT-blue.svg)

</p>

</div>

## About Gechoplate

Gechoplate is a simple MVC boilerplate to design Rest APIs in Golang. It provides several basic features like a JWT authentication system, a complete configuration file system and a highly optimized HTTP router.

This boilerplate has been designed to be used by everyone and especially for developers new to the Go language, feel free to use it in your personnal or school projects !

### Features

* Highly optimized HTTP router & built-in Middleware with [Echo Framework](https://github.com/labstack/echo)
* Easy and clean configuration file system with [Viper](https://github.com/spf13/viper)
* JWT Authentification system with [jwt-go](https://github.com/dgrijalva/jwt-go)
* MySQL connexion support with [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)
* All in one [Docker compose](https://docs.docker.com/compose/install/) with a Go build and a MySQL server
* Friendly automation tool for project management with *Make*
* Preconfigured helpers like a HTTP parameter validator and a JSON response setter

## Requirement

Gechoplate comes with an all-in-one docker-compose including a Go build, a MySQL server and a SQL dump initialization system. 

To use the boilerplate in an optimal way, docker is required.

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/install/)

## Use this boilerplate

With [GitHub CLI](https://cli.github.com/) :

```sh
$ gh repo create [name] --template Akecel/gechoplate
```

Or use the button **[Use this template](https://github.com/Akecel/gechoplate/generate)** above

You can also clone this repository :

```sh
$ git clone https://github.com/Akecel/gechoplate.git
```

## Configuration

Gechoplate is using [Viper](https://github.com/spf13/viper) to provide a complete configuration file system.

You will need to set your environement file to run it :

```bash
$ cp .env.example .env
```

Docker-compose and Gechoplate use the same environment variables.

This way you have only one file to configure :

```bash
APP_URL=http:localhost:1323
APP_NAME=Gechoplate
DB_HOST=db
DB_PORT=3306
DB_NAME=database
DB_USER=user
DB_PASSWORD=password
```

## Usage

### Makefile

Gechoplate use a Makefile to manage all commands of the project.

You can display the list of commands with :

```sh
$ make help
Usage: make <command>

Commands:
  help                   Provides help information on available commands
  compose/build          Build all Docker images of the project
  compose/up             Start all containers (in the background)
  compose/down           Stops and deletes containers.
  compose/purge          Stops and deletes containers, volumes, images and networks.
  compose/rebuild        Rebuild the project
  urls                   Get project's URL
```

### Run application

Start the project with :

```sh
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

You can now access the api: [http://localhost:1323/](http://localhost:1323/).

## Test

Gechoplate has a *tests* directory containing all the test files for each controller.\
The test files are generated using [gotests](https://github.com/cweill/gotests).

Simply use the test go command to run your tests :

```sh
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

This project is licensed under the [MIT License](https://opensource.org/licenses)  - see the [LICENSE.md](https://github.com/Akecel/gechoplate/blob/master/LICENCE) file for details.

## Show your support

Give a ⭐️ if this project helped you!