FROM golang:alpine as builder

#For installing go dependencies
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app
COPY . .

#Fetching go dependencies
RUN go mod download

#Building the executable for linux only
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/levels

FROM scratch

COPY --from=builder /go/bin/levels /go/bin/levels
EXPOSE 9091

ENTRYPOINT ["/go/bin/levels"]