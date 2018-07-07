FROM golang:1.10.3 AS builder

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/nclandrei/bowlingo
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /bowlingo cmd/main.go

FROM scratch
COPY --from=builder /bowlingo ./
EXPOSE 80
CMD ["./bowlingo", "-port=80"]