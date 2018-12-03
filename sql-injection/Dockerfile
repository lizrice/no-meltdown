FROM golang
RUN go get -u github.com/securego/gosec/cmd/gosec/...
WORKDIR /go/src/github.com/justincormack/dockercon-sql-injection
COPY . .
# RUN gosec ./...
RUN go install .
EXPOSE 80
ENTRYPOINT ["/go/bin/dockercon-sql-injection"]
