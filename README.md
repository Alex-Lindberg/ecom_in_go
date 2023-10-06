# Ecommerce backend in Go

Testing out programming in Go. 

This project consists of a small ecommerce backend service.
To run I usually use:

```shell
docker-compose up -d && docker-compose logs -f app
```

### Package structure

Borrowed a structure form somewhere, using as a rough guidline.

[Directory/Package]: [Functionality]
- api: Delivery layer
    - Externals, HTTP etc.
    - Parsing
- cmd: Distribution
    - Main
    - Parse configs
- models: Data access layer (A)
    - Shared objects
    - "IAM"
    - Transforms
- storage: Data access layer (B)
    - Data access interfaces (pg DAO)
- usecases: Buisness logic
    - Interfaces for operations over entities (features)
    - mocking
    - extensibility
- tests: Self explanatory
- Makefile

