# Build stage
FROM golang:1.23 AS builder
WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/ads ./cmd/server/main.go

# Final stage
# FROM scratch
# COPY --from=builder /bin/ads /bin/ads
CMD ["/bin/ads"]
