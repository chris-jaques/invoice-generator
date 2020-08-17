FROM golang:1.14

# Build Dependencies
RUN go get -u gopkg.in/yaml.v2 github.com/shurcooL/graphql

WORKDIR /go/src/invgen
COPY ./invgen .
COPY ./config.yaml /root/invgen.conf

# RUN go get -d -v ./... 

# RUN go install -v ./...
# RUN go build -o /usr/local/bin/geninv geninv.go
