FROM golang:1.25

# Install Chromium (for chromedp PDF generation)
RUN apt-get update && apt-get install -y --no-install-recommends \
    chromium \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

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
