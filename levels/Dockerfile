FROM golang:alpine as builder

#For installing go dependencies
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN adduser -D -g '' appuser

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#Building the executable for linux only
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o users

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/users  users
USER appuser

EXPOSE 9090


ENTRYPOINT ["./users"]