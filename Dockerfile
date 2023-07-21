FROM golang:1.19
# Define build env
ENV GOOS linux
ENV CGO_ENABLED 1
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build -o app

COPY ./config.json /config/config.json

# Exec built binary
CMD ./app run --config=/config/config.json