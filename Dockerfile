FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
#ADD go.sum .
RUN go mod download
COPY ./cmd ./cmd
#COPY ./web/ ./web/
RUN go build -o app ./cmd/main/app.go

FROM alpine
#RUN apk update --no-cache && apk add --no-cache ca-certificates
WORKDIR /src
EXPOSE 80
COPY --from=builder /build/app app
#COPY --from=builder /build/web web
COPY ./web/ ./web
CMD ["./app"]
