FROM golang:alpine as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o cmd/main.go


FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/cmd/ /app/

WORKDIR /app

EXPOSE 3000

CMD ["./cmd/"]