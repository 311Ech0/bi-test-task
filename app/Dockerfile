FROM golang:latest

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64   
COPY *.go .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
