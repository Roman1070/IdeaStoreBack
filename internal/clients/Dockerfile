FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /var/backend

COPY internal/clients/auth.go clients/auth.go
COPY internal/clients/boards.go clients/boards.go
COPY internal/clients/chats.go clients/chats.go
COPY internal/clients/ideas.go clients/ideas.go
COPY internal/clients/comments.go clients/comments.go
COPY internal/clients/profiles.go clients/profiles.go
COPY internal/clients/main.go clients/main.go
COPY . .

RUN go mod tidy
RUN go build -o main clients/main.go clients/auth.go clients/boards.go clients/chats.go clients/ideas.go clients/comments.go clients/profiles.go

FROM alpine:edge AS prod

RUN apk add bash

COPY --from=build /var/backend/main /app/main
COPY --from=build /var/backend/.env /app/.env

RUN mkdir -p /app/files
COPY files /app/files

RUN mkdir -p /app/configs
COPY configs app/configs

WORKDIR /app
EXPOSE 8000

ENTRYPOINT ["./main"]