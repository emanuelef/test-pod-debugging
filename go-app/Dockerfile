FROM golang:1.21.6-alpine AS builder
WORKDIR /usr/src/app/
COPY main.go .
COPY go.mod .
COPY go.sum .
RUN go build -o /go/bin/go-app/

FROM alpine:latest AS runner
WORKDIR /usr/src/app/
COPY --from=builder /go/bin/go-app/ .
EXPOSE 8080
ENTRYPOINT [ "./go-app" ]