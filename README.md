# Cook Book  - API

### Overview
The purpose of the project is to provide an API to manage recipes and ingredients in order to filter for recipes by ingredients.

###  Documentation
[OpenAPI](./docs/openapi.yaml)

### Getting started
The API is composed of two parts, the first is an `REST API` that helps to manage the searches and storage related to recipes and the second is a `CLI` helps to fill the recipe storage.

### How to use the CLI
###### Compile
Follow the instructions below to compile the binary file for the `CLI` in the `/build` directory
```shell
make cli
```
###### See how to use
```shell
./build/cli
```

### How to use the REST API
###### Compile
Follow the instructions below to compile the binary file for the `REST API` in the `/build` directory
```shell
make http
```
###### Load environment variables
```shell
export $(grep -v ^# .env.example)
```
###### Start the server
```shell
./build/http
```

### Scripts
This project contains some bash scripts to help to make some operations like compile.
[See bash scripts](./scripts)

### Architecture
###### Pattern (also style)
This project implements architecture pattern [Ports and Adapters](https://alistair.cockburn.us/hexagonal-architecture)
###### Decisions
**Vertical Slicing**

Interpreting what [Vertical Slicing](https://en.wikipedia.org/wiki/Vertical_slice) says, I decided to make one package per feature and put a little of each layer in each package.

**Go Project Layout Standard**

I decided to follow the [Go Project Layout Standard](https://github.com/golang-standards/project-layout).

**Compile only what is required**

According to the theory of hexagonal architecture, it is possible to have *n* adapters for different external signals (http, gRPC, command line).

*For example* one use case is to create a recipe, but the instruction comes from a http request or a kafka message.

So I decided to compile a binary to handle each signal.

###### Package tree
```
.
├── cmd
│   └── {binary}
├── internal
│   ├── app
│   │   └── {feature}
│   │       ├── business (Use cases, rules, data models and ports)
│   │       └── infrastructure
│   │           ├── input  (Everything related to "driving" adapters)
│   │           └── output (Everything related to "driven" adapters)
│   └── container (DI container)
└── pkg (Public and global code, potencially libraries)
```