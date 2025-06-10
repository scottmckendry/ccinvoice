FROM golang:1.24

RUN apt-get update && apt-get install wkhtmltopdf -y

WORKDIR /app
COPY . /app/
RUN go mod download
RUN go mod verify
RUN go install github.com/air-verse/air@latest
CMD air -c .air.toml
