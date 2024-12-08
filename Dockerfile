FROM golang:1.21.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["CompileDaemon", "--build=go build -o main .", "--command=./main"]


EXPOSE 8080

CMD ["./main"]
