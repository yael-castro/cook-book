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
make server
```
###### Load environment variables
```shell
export $(grep -v ^# .env.example)
```
###### Start the server
```shell
./build/server
```

### Scripts
This project contains some bash scripts to help to make some operations like compile.

[See bash scripts](./scripts)

### Architecture pattern
This project implements architecture pattern [ports and adapters](https://alistair.cockburn.us/hexagonal-architecture)
```
.
├── cmd
│   └── {application}
│       └── container (di container)
│
├── internal
│   └── {topic}
│       ├── business
│       │   ├── logic  (use cases)
│       │   ├── model  (data models)
│       │   └── port   (ports)
│       └── infrastructure
│           ├── input  (everything related to "driving" adapters)
│           └── output (everything related to "driven" adapters)
│
└── pkg (global and public code, potentially libraries)
```