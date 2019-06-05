FROM golang:alpine as builder

#For installing go dependencies
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app
COPY . .

#Fetching go dependencies
RUN go mod download

#Building the executable for linux only
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/players

FROM scratch

COPY --from=builder /go/bin/players /go/bin/players
EXPOSE 9090

ENTRYPOINT ["/go/bin/players"]