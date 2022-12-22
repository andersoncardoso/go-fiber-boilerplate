# Go API boilerplate

Uses go, gorm (postgresql) and fiber

The structure is:

- routers: handle http requests, and call to the service layer
- services: services contains the bussiness logic, it's also the only layer that should access the models
- models: gorm models definitions

## Development

### Start the application

```bash
go install github.com/mitranim/gow@latest
gow run app.go

# flags
go run app.go -prod
go run app.go -port=:3000
```

```
# Clean packages
make clean-packages

# Generate go.mod & go.sum files
make requirements

# Compile project
make build

# Run the project in prefork mode
make run

# Run in development mode
makr run-dev
```
