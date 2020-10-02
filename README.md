<h1 align="center">Go-Boilerplate for Rest APIs 🚀</h1>

<p align="center">
  <a href="https://github.com/labstack/echo" target="_blank">
    <img alt="echo" src="https://cdn.labstack.com/images/echo-logo.svg" width="40%"/>
  </a>
</p>

<p align="center">
  <img alt="Version" src="https://img.shields.io/badge/version-0.1-blue.svg?cacheSeconds=2592000" />

  <a href="https://godoc.org/github.com/Akecel/api-goilerplate" target="_blank">
    <img alt="godoc" src="https://godoc.org/github.com//Akecel/api-goilerplate?status.svg" />
  </a>

  <a href="https://goreportcard.com/report/github.com/Akecel/api-goilerplate" target="_blank">
    <img alt="goReport" src="https://goreportcard.com/badge/github.com/Akecel/api-goilerplate" :>
  </a>

  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

## Introduction

Goilerplate is a [Echo v4](https://github.com/labstack/echo) boilerplate  to design simple Rest APIs in Golang. Using the MVC design pattern, it provides a JWT authentication system, a configuration file system thanks to [Viper](https://github.com/spf13/viper). It also contain a complete routing system and pre-configured helpers.

Feel free to check the [Echo documentation](https://echo.labstack.com/guide) too use this template properly.

## Requirement

This template contains a docker compose with Go and MySQL and a SQL dump migration system.

* [Docker](https://www.docker.com/)

* [Docker Compose](https://docs.docker.com/compose/install/)

## Use this boilerplate

With [GitHub CLI](https://cli.github.com/) :

```sh
$ gh repo create [name] --template Akecel/api-goilerplate
```

Or use the button **[Use this template](https://github.com/Akecel/api-goilerplate/generate)** above

## Configuration

This template is using [Viper](https://github.com/spf13/viper) to provide a configuration file system, you will need to set your configuration file to use Goilerplate :

```bash
$ cp config.example.yml config.yml
```

```bash
//config.yml

APP_NAME: api-goilerplate
DB_HOST: db
DB_PORT: 3306
DB_NAME: database
DB_USER: user
DB_PASSWORD: password
```

## Usage

```sh
TODO
```

## Test

```sh
TODO
```

## Contributing

Please read [CONTRIBUTING.md](https://github.com/Akecel/api-goilerplate/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* [**Akecel**](https://github.com/Akecel) - *Template author*
* [**Hugo-T**](https://github.com/T-Hugo) - *Development environment author*

See also the list of [contributors](https://github.com/Akecel/api-goilerplate/graphs/contributors) who participated in this project.

## Licence

This project is licensed under the [MIT License](https://opensource.org/licenses)  - see the [LICENSE.md](https://github.com/Akecel/api-goilerplate/blob/master/LICENCE) file for details.


## Show your support

Give a ⭐️ if this project helped you!