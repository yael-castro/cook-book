# Cook Book  - API
The purpose of the project is to provide an API to manage recipes and ingredients in order to search for recipes by ingredients.

### More about the project
This project is compose of two parts very important.
The first part is the `REST API` that expose the endpoints to query and manage recipes
and the second is the `CLI` to manage internal operations and help to test the search engine.

### Getting started

###### How to run

```shell
make run
```

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
