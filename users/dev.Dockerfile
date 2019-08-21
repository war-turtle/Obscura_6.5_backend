FROM golang:alpine as builder

RUN apk update && apk add --no-cache git npm ca-certificates && update-ca-certificates
RUN npm i -g nodemon

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o users

EXPOSE 9090

CMD [ "nodemon", "-e", "go", "--signal", "SIGINT", "--exec", "go", "run", "*.go" ]