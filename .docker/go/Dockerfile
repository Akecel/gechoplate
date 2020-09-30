FROM golang:1.14.1-alpine AS builder
LABEL stage=builder

WORKDIR /go/src

COPY go.mod go.sum ./
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build \
    -a \
    -trimpath \
    -ldflags '-s -w -extldflags "-static"' \
    -tags 'netgo osusergo static_build' \
    -o /go/bin/app

# Copying the binary into the final image
FROM alpine
COPY --from=builder /go/bin/app .
ENTRYPOINT ["./app"]