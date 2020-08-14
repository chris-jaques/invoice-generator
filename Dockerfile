FROM golang:1.14

# Build Dependencies
RUN go get gopkg.in/yaml.v2

WORKDIR /go/src/invgen
COPY ./invgen .
COPY ./config.yaml /root/invgen.conf

# RUN go get -d -v ./... 

# RUN go install -v ./...
# RUN go build -o /usr/local/bin/geninv geninv.go
