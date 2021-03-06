# Golang Sandbox

A sandbox/playgorund to learn and explore the golang language.

This repo covers;

- running a gin api
- loops
- variables
- conditionals
- operands and pointers
- arrays
- slices
- structs
- strings module
- errors module
- time module
- math module
- environment
- json web tokens (jwt)

### Running

`go run .`

### Building

`go build`

### Docker/Postgres commands

```
"start": "docker run --name some-pg -p 3008:5432 -e POSTGRES_PASSWORD=secret -d postgres",
"kill": "docker kill some-pg && docker rm some-pg",
"restart": "docker restart some-pg",
"console": "docker exec -it some-pg psql -U postgres",

```
