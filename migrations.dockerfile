FROM golang:alpine as builder

WORKDIR /app
ADD cmd/ /app/cmd
ADD pkg/ /app/pkg
COPY go.mod /app
COPY go.sum /app

RUN mkdir /build

RUN CGO_ENABLED=0 GOOS=linux GIN_MODE=release go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /build ./cmd/migrations

FROM scratch

WORKDIR /app
COPY --from=builder /build /app

ADD data/ /app/data

CMD ["/app/migrations"] 