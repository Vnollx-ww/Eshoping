FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/main cmd/api/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/main

COPY ../config/ /app/config/

COPY ../web/ /app/web/

EXPOSE 8889
CMD ["./main"]