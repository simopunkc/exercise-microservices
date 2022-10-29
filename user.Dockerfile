##
## Build
##
FROM golang:alpine AS build

WORKDIR /app

COPY ./user-service /app

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o user cmd/api/main.go

##
## Deploy
##
FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=build /app/env.example /app/.env
COPY --from=build /app/user /app/user

USER nonroot:nonroot

CMD ["./user"]