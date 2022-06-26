FROM golang:1.18

# Copy source files
RUN mkdir /app
COPY . /app

# Build binary
WORKDIR /app
RUN go build -o /webapp .

# Run binary
ENTRYPOINT ["/webapp"]