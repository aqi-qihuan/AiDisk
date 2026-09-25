FROM golang:1.27.1-alpine AS builder

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /agent ./cmd/agent

FROM alpine:3.22
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /agent .
EXPOSE 8000
CMD ["./agent"]
