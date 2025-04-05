FROM golang:1.23

WORKDIR /app
COPY . /app/
RUN go mod tidy
RUN go build -o main .
RUN apt-get update && apt-get install wkhtmltopdf -y
EXPOSE 3000
ENTRYPOINT ["./main"]
