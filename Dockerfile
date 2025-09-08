FROM golang:1.22-bullseye AS builder
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /src

COPY go.mod go.sum ./
COPY vendor/ ./vendor
COPY . .

WORKDIR /src/cmd/server
ENV GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o /server

FROM gcr.io/distroless/static-debian12
COPY --from=builder /server /server
ENV PORT=8080
EXPOSE 8080
CMD ["/server"]
