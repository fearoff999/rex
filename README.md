# rex
Small tool for remote command execution via hook

## Installation

Download binary from [releases page](https://github.com/fearoff999/rex/releases)

Make binary executable
    `sudo chown $USER:$USER rex`

## Usage

1. Simply define 3 variables at `.env.rex` file (described below, example is provided at repo `.env.rex.example`).
2. Run binary with `./rex &` or `nohup ./rex &`
3. Trigger command execution with GET-request `http://#host#:#REX_PORT#/`
4. Wait till end of execution and see stdout/stderr of command at response
5. Profit

## Environment variables

`REX_USER` - username used for basic auth

`REX_PASSWORD` - password used for basic auth

`REX_COMMAND` - command that should be executed (e.g: `docker-compose restart webserver`, `date`, ...)

`REX_PORT` - http webserver port

## Healthcheck

`http://#host#:#REX_PORT#/health` can be used as healthcheck of rex service (responses with `200` if service is up)