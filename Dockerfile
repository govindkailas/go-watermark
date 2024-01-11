FROM golang:1.21-alpine AS build
COPY  . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o go-watermark-app

FROM alpine:latest
COPY --from=build /app/go-watermark-app .
EXPOSE 8080
ENTRYPOINT [ "./go-watermark-app" ]

RUN apk add --no-cache bash curl