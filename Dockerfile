FROM golang:1.19-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN go build -o server .

FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server"]