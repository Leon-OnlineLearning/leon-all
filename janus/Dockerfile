FROM golang:alpine AS builder

WORKDIR /code
COPY ./go_server .

RUN CGO_ENABLED=0 go build main.go


FROM canyan/janus-gateway:latest

WORKDIR /code

COPY --from=builder /code/main /code/main
COPY start.sh start.sh

EXPOSE 6111
CMD ["./start.sh"]


