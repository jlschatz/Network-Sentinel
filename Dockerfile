FROM golang:1.12-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/daApp

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

RUN go build -o ./daApp .

FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/daApp/daApp /app/daApp

EXPOSE 8080

CMD ["/app/daApp"]