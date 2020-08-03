FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod vendor

RUN go build -o main .

WORKDIR /dist
RUN cp /build/.env .
RUN cp /build/main .

EXPOSE 8087

FROM scratch

COPY --from=builder /dist/main /
COPY --from=builder /dist/.env /
EXPOSE 8087

ENTRYPOINT ["/main"]
