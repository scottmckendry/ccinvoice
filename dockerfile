FROM golang:1.21

WORKDIR /app
COPY . /app/
RUN go build -o main .
RUN apt-get update && apt-get install wkhtmltopdf -y
EXPOSE 3000
ENTRYPOINT ["./main"]
