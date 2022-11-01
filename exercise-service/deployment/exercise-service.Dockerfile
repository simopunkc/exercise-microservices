##
## Build
##
FROM golang:alpine AS build

WORKDIR /app

COPY ./exercise-service /app

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o exercise cmd/api/main.go

##
## Deploy
##
FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=build /app/config/env.example /app/config/.env
COPY --from=build /app/exercise /app/exercise

USER nonroot:nonroot

CMD ["./exercise"]