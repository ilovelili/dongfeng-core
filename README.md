# Core

Core is the microservice architecture interacts with frontend.

## Dependencies

- [Go 1.10](https://golang.org/) aka Golang
- [Micro](https://micro.mu/) Microservice toolkit written in Go
- [NATS](http://nats.io/) High performance message broker
- [Redis](https://redis.io/) Cache server meanwhile saves session
- [MySQL](https://www.mysql.com/) RDBMS to save core business data
- [Auth0](https://auth0.com/) Robust authentication service
- [Consul](https://www.consul.io/) Local service discovery (not necessary for production)
- [Docker](https://www.docker.com/) Containerization platform

## Go packages

### Public packages

- Please check `Gopkg.lock`

## Spinup

Step 1. Start local services

```bash
cd local && run.sh
```

Step 2. Start core

```bash
cd core && run.sh
```

[Git Bash](https://git-scm.com/downloads) for Windows

## Endpoints (TBD)