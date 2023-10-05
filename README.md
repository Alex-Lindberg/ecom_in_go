# Ecommerce backend in Go

### Goal package structure

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

