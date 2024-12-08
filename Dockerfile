FROM golang:1.21.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . .

RUN go build -o main .


EXPOSE 8080

CMD ["./main"]

# CMD ["CompileDaemon", "--build=go build -o main .", "--command=./main"]