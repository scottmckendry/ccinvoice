FROM golang:1.25

# Install Chromium - necessary for chromedp and pdf generation
RUN apt-get update && apt-get install -y --no-install-recommends \
    chromium \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY . /app/
RUN go mod download
RUN go mod verify
RUN go install github.com/air-verse/air@latest
CMD air -c .air.toml
