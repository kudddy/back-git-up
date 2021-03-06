FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    db_name=db \
    db_pass=pass \
    db_user=user \
    db_type=postgres \
    db_host=docker.for.mac.host.internal \
    db_port=5434 \
    memcache_addres=docker.for.mac.host.internal:11211 \
    frontend_adress=http://127.0.0.1:3000

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
COPY .env .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/dist/main"]