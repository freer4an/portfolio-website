# Build stage
FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app
COPY . .

RUN go build -o main cmd/*.go

# Final stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/configs configs
COPY --from=builder /app/front front

EXPOSE 8000
CMD [ "/app/main" ]