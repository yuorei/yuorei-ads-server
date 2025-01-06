FROM golang:1.23
WORKDIR /src
COPY . .

RUN go build -o /bin/ads ./cmd/server/main.go

FROM scratch
COPY --from=0 /bin/ads /bin/ads
CMD ["/bin/ads"]
