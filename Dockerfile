# docker build -t my-golang-app .
# docker run --rm -v "$PWD":/usr/src/curl_elk -w /usr/src/curl_elk my-golang-app go build -v

FROM golang:1.14

WORKDIR /go/src/cats_request
COPY . .

RUN go get -d -v ./...
#RUN go install -v ./...

CMD ["cats_request"]
