FROM golang:1.25

# Install Chrome - necessary for chromedp and pdf generation
RUN apt-get update && apt-get install -y \
    wget \
    gnupg \
    && wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list \
    && apt-get update \
    && apt-get install -y google-chrome-stable \
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
