# Go API Server for openapi

このサービスはスキーのサポートをするため、スキー場、スキー場の天気、スキー場のホテル、ホテルの予約などの機能を提供する。

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 1.0
- Build date: 2024-01-16T06:31:36.418090726Z[Etc/UTC]


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t openapi .
```

Once image is built use
```
docker run --rm -it openapi
```

Generate code from OpenAPI defination
```
podman run --rm -v ${PWD}:/local docker.io/openapitools/openapi-generator-cli \
       generate \
       -i /local/snowman-booking.yml \
       -g go-server \
       -o /local/generated
```
