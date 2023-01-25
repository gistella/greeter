FROM golang:1.19
WORKDIR /go/src/github.com/gistella/greeter
COPY src/* ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM registry.access.redhat.com/ubi8/ubi:8.7
WORKDIR /root/
COPY --from=0 /go/src/github.com/gistella/greeter/app ./
CMD ["./app"]
