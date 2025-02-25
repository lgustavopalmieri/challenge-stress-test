FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /main

ENTRYPOINT ["/main"]