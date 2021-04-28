FROM golang:1.16-alpine as builder

WORKDIR /go/src/github.com/badis/hackathon

ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
COPY . .
RUN go install -v ./...

FROM scratch  
COPY --from=builder /go/bin/hackathon /go/bin/hackathon
ENTRYPOINT ["/go/bin/hackathon"]