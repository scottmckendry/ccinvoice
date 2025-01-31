FROM golang:1.23

WORKDIR /app
COPY . /app/
RUN go mod download
RUN go mod verify
RUN go install github.com/air-verse/air@latest
RUN apt-get update && apt-get install wkhtmltopdf -y
CMD air -c .air.toml
