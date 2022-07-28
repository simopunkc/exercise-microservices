# User Service
Service for get user

## Dependencies
- mysql
- golang 1.17

## API
### Public
- login
- register

### Internal
- get by user_id

## Project Structure

```
├── cmd
|   ├── api
│   │   └── main.go
│   └── [other service]
│       └── main.go
├── internal
|   ├── app
|   |   ├── domain
|   |   |   └── domain.go
|   |   ├── datahase
|   |   |   └── mysql.go
|   |   ├── user
|   |   |   ├── repository
|   |   |   |   └── repository.go
|   |   |   ├── handler
|   |   |   |   └── handler.go
|   |   |   ├── service/usecase
|   |   |   |   └── service.go
|   |   ├── [other domain]
|   |   |   ├── repository
|   |   |   |   └── repository.go
|   |   |   ├── handler
|   |   |   |   └── handler.go
|   |   |   ├── service/usecase
|   |   |   |   └── service.go
|   ├── pkg
|   |   ├── middleware
|   |   |   └── basic_auth.go
|   |   ├── [other package]
|   |   |   └── [other package].go
```