<h1 align="center"> Let's use Gechoplate 🚀</h1>

<p align="center">
  <a href="https://github.com/labstack/echo" target="_blank">
    <img alt="echo" src="https://cdn.labstack.com/images/echo-logo.svg" width="40%"/>
  </a>
</p>

<p align="center">
  <img alt="Version" src="https://img.shields.io/badge/version-0.1-blue.svg?cacheSeconds=2592000" />

  <a href="https://godoc.org/github.com/Akecel/gechoplate" target="_blank">
    <img alt="godoc" src="https://godoc.org/github.com//Akecel/gechoplate?status.svg" />
  </a>

  <a href="https://goreportcard.com/report/github.com/Akecel/gechoplate" target="_blank">
    <img alt="goReport" src="https://goreportcard.com/badge/github.com/Akecel/gechoplate" />
  </a>

  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-blue.svg" />
  </a>
</p>

## About Gechoplate

Gechoplate is a [Echo v4](https://github.com/labstack/echo) boilerplate  to design simple Rest APIs in Golang. Using the MVC design pattern, it provides a JWT authentication system, a configuration file system thanks to [Viper](https://github.com/spf13/viper). It also contain a complete routing system and pre-configured helpers.

Feel free to check the [Echo documentation](https://echo.labstack.com/guide) too use this template properly.

## Requirement

Gechoplate contains a docker compose with Go and MySQL and a SQL dump migration system.

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

Gechoplate is using [Viper](https://github.com/spf13/viper) to provide a configuration file system, you will need to set your environement file to use it :

```bash
$ cp .env.example .env
```

Docker-compose and Gechoplate use the same environment variables, so you have only one file to configure :

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

Gechoplate use a Makefile to manage all commands of the project, you can display the list of commands with :

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

[Echo](https://github.com/labstack/echo) provide a complete test system for your application, for more details on the testing process of echo, check the [documentation](https://echo.labstack.com/guide/testing).

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