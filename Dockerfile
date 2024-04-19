FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./bin/main ./cmd/*.go


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/bin/main .
EXPOSE 8080
ENTRYPOINT ["./bin/main"]