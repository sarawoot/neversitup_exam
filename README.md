```
Project
├── cmd
│   ├── migration
│   │   └── main.go
│   └── server
│       └── main.go
├── env
│   ├── dev.json
│   └── local.env
├── gen
│   ├── protobuf_service
│       └── v1
│           ├── protobuf_service.pb.go
│           └── protobuf_service_grpc.pb.go
├── go.mod
├── go.sum
├── internal
│   ├── adapter
│   │   ├── callexternal_service.go
│   ├── constant
│   │   └── status.go
│   ├── core
│   │   ├── domain
│   │   │   ├── model.go
│   │   ├── port
│   │   │   ├── interface{service,repository}.go
│   │   └── service
│   │       ├── service_name.go
│   ├── handler
│   │   ├── grpc_handler.go
│   ├── infrastructure
│   │   └── grpc_server.go
│   ├── property
│   │   └── property.go
│   ├── repository
│   │   ├── repository.go
├── proto-hub{git submodule}
├── migrations
        └── 01_create_tables.sql
```