FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Install air
RUN go install github.com/cosmtrek/air@latest

RUN go build -o /go/bin/app

expose 8080

CMD ["/go/bin/air"]