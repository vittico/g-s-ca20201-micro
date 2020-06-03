FROM golang:1.14
WORKDIR /go/src/github.com/vittico/g-s-ca20201-micro
COPY . .
RUN pwd && ls -lkah
RUN go mod download && go mod vendor && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM alpine:latest
WORKDIR /opt/gsca20201-micro
COPY --from=0 /go/src/github.com/vittico/g-s-ca20201-micro/g-s-ca20201-micro .
EXPOSE 10000/tcp
CMD [ "./g-s-ca20201-micro" ]