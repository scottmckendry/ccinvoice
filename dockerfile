FROM golang:1.24

RUN apt-get update && apt-get install wkhtmltopdf -y

# Create non-root user
RUN useradd -u 1000 -m ccinvoice

WORKDIR /app
COPY go.mod go.sum ./
COPY migrations/ ./migrations/
COPY public/ ./public/
COPY views/ ./views/
COPY *.go ./

RUN go mod tidy
RUN go build -o main .

# Clean up
RUN rm *.go
RUN rm go.*

RUN chown -R ccinvoice:ccinvoice /app
USER ccinvoice
EXPOSE 3000
ENTRYPOINT ["./main"]
