FROM golang:1.8
WORKDIR /go/src/github.com/ignoshi/snippets
COPY . /go/src/github.com/ignoshi/snippets
RUN go get
EXPOSE 8000
CMD ["go", "run", "main.go"]
