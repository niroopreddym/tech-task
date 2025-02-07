FROM golang:1.15.0 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

RUN export GO111MODULE=on

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .
# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /app/main .
# Expose port 9000 to the outside world
EXPOSE 9294
# executable
ENTRYPOINT [ "./main" ]