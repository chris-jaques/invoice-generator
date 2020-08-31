FROM golang:1.14

# Build Dependencies
RUN go get -u gopkg.in/yaml.v2 \
    github.com/shurcooL/graphql \
    golang.org/x/oauth2 \
    golang.org/x/net/context

WORKDIR /go/src/invgen
COPY ./invgen .
COPY ./config.yaml /root/invgen.conf

ENV ENVIRONMENT=production
# RUN go get -d -v ./... 

# RUN go install -v ./...
# RUN go build -o /usr/local/bin/geninv geninv.go
