FROM golang:1.10-alpine

EXPOSE 8080

WORKDIR /go/src/github.com/Oscar170/i-rol-back/user-api

# Dependency management
RUN apk --no-cache add curl git && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy files.
COPY . .

# Install dependency
RUN dep ensure

RUN go build

CMD [ "./user-api" ]
