FROM golang:1.14
WORKDIR /app
COPY . .
RUN go get -d -v ./...
EXPOSE 8000
CMD ["go" , "run", "main.go"]