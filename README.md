# Cook Book  - API
The purpose of the project is to provide an API to manage recipes and memory in order to search for recipes by memory.

### Getting started
This project is compose of multiple parts very important.

1. `REST API` that expose the endpoints to query and manage recipes.
2. `CLI` to manage internal operations and help to test the search engine.

### How to build the binary files
Follow the instructions below to compile the binaries in the **/build** directory
###### CLI
```shell
make cli
```
###### REST API
```shell
make server
```

### Architecture pattern
This project implements architecture pattern "ports and adapters"
```

```