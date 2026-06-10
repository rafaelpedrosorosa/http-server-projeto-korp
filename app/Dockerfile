FROM golang:1.23-alpine AS build

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /http-server-projeto-korp .

FROM alpine:3.20

RUN adduser -D -H appuser
USER appuser

COPY --from=build /http-server-projeto-korp /http-server-projeto-korp

EXPOSE 8080
ENTRYPOINT ["/http-server-projeto-korp"]
