# ---- Build Stage: Compile the Go Application ----
FROM golang:1.24.1-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# ---- Production Stage: Create a Minimal and Secure Image ----
FROM alpine:3.20.1 AS prod
WORKDIR /app

RUN adduser -D appuser
USER appuser

COPY --from=build /app/main /app/main

COPY --from=build /app/internal/challenges/challenges.json /app/internal/challenges/challenges.json

ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}

ENTRYPOINT ["/app/main"]
