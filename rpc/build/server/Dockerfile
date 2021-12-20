FROM golang:1.16 AS build

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

COPY . .
RUN go build -ldflags="-s -w" -o /app/grpc-serve cmd/server/main.go


FROM alpine as prod

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=build /app/grpc-serve /app/grpc-serve

CMD ["./grpc-serve"]
