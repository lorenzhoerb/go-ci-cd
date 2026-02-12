FROM golang:1.25.5 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o app

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]