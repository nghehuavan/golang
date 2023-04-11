<bounded-context>
├──config
│   ├── app_config.json
│   ├── config.go
│   └── schema.json
├── clients
│   └── <other-microservice>
│       └── client.go
├── constants
│   └── constants.go
├── controller
│   ├── <foo>_controller.go
│   └── ...
├── db
│   └── database.go
├── error
│   ├── api_error.go
│   └── custom_error.go
├── mocks
│   ├── <foo-repository>_repository_mock.go
│   ├── <foo-service>_service_mock.go
│   ├── <other-microservice>_client_mock.go
│   └── ...
├── models
│   ├── request
│   │   ├── <api-endpoint>_request.go
│   │   └── ...
│   ├── response
│   │   ├── <api-endpoint>_response.go
│   │   └── ...
│   ├── <foo-models>.go
│   ├── <bar-models>.go
│   └── ...
├── repository
│   ├── <foo-respository>_repository.go
│   └── ...
├── router
│   └── routes.go
├── service
│   ├── <foo-service>_service.go
│   └── ...
└── utils
    └── ...