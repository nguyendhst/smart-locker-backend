FROM golang:1.20
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod tidy
RUN go build -o main .
CMD ["/app/main"]
