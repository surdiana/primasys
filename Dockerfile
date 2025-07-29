FROM golang:1.22-alpine AS build
WORKDIR /app
COPY main.go .
RUN go build -o server main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 8080
ENTRYPOINT ["./server"]