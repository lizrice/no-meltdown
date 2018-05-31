FROM golang
RUN go get -u github.com/GoASTScanner/gas/cmd/gas/...
WORKDIR /go/src/github.com/justincormack/dockercon-sql-injection
COPY . .
# RUN gas ./...
RUN go install .
EXPOSE 80
ENTRYPOINT ["/go/bin/dockercon-sql-injection"]
