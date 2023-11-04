FROM golang:1.21

WORKDIR /app
COPY . /app/
RUN go build -o main .
EXPOSE 3000
ENTRYPOINT ["./main"]
