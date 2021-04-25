FROM golang:1.16-alpine

WORKDIR /go/src/github.com/badis/hackathon
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["hackathon"]
