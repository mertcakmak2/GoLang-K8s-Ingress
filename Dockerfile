FROM golang:1.15-alpine as builder
WORKDIR /go/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app main.go
FROM gcr.io/distroless/base
COPY --from=builder /go/app/ .
CMD ["/app"]

# FROM golang:1.15-alpine AS builder 
# WORKDIR /build
# ADD go.mod .
# COPY . .
# RUN go build -o main main.go
# FROM alpine
# WORKDIR /build
# COPY --from=builder /build/main /build/main
# CMD [". /main"]