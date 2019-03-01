FROM golang:latest as builder
WORKDIR /go/src/github.com/ivankg/go-web-hello-world
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
EXPOSE 80
COPY --from=builder /go/src/github.com/ivankg/go-web-hello-world/main .
ENTRYPOINT ["/main"]
