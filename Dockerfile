FROM golang
COPY meliScanner /opt/meliScanner
WORKDIR /opt/meliScanner
RUN go get -d ./...
RUN go build adapter/http/main.go
ENTRYPOINT [ "./main" ]