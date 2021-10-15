FROM golang:1.17

WORKDIR /service

COPY . .
RUN make build && go build -o service cmd/main.go
CMD ["./service"]