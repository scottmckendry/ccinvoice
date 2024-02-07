FROM golang:1.22

WORKDIR /app
COPY . /app/
RUN go mod download
RUN go mod verify
RUN go install github.com/cosmtrek/air@latest
RUN apt-get update && apt-get install wkhtmltopdf -y
CMD air -c .air.toml
