FROM golang:1.14-alpine AS builder
RUN apk add git
WORKDIR /app
COPY . .
RUN go get
RUN go build

FROM alpine:3.12
WORKDIR /app
COPY --from=builder /app/assets /app/assets
COPY --from=builder /app/img-api /app/
ENTRYPOINT ["./img-api", "-h", "0.0.0.0"]