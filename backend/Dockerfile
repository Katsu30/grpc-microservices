FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download
RUN go get github.com/cosmtrek/air@v1.29.0

COPY . .

RUN go build -o main .

CMD sh -c "air -c .air.toml && go run main.go"
