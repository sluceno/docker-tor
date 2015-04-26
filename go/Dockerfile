FROM golang:1.4.2

# ENTRYPOINT /go/src

COPY ./src/github.com/sluceno/go-tor /go/src/github.com/sluceno/go-tor

RUN export PATH=$PATH:$GOPATH/bin

RUN go get -v github.com/sluceno/go-tor

RUN go install github.com/sluceno/go-tor

# CMD ["go", "run", "/go/src/github.com/sluceno/go-tor/main.go"]
