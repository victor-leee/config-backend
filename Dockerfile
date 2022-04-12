FROM golang:latest
WORKDIR /go/src/github.com/victor-leee/config-backend
COPY . .
RUN go build -o main cmd/server/main.go
CMD ["./main"]