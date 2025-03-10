# Build stage
FROM --platform=$BUILDPLATFORM golang:1.23.3-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o endpointlab main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/endpointlab .

# Copy images to the container
COPY templates/images /app/templates/images

# Set the image path environment variable
ENV IMAGE_PATH=/app/templates/images

EXPOSE 8080
ENTRYPOINT ["/app/endpointlab"]
