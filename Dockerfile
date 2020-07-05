FROM golang:1.12.0-alpine3.9
# Adding the libraries neccessary
RUN apk update && apk add
RUN apk add --no-cache git
RUN go get -u github.com/gorilla/mux
# create a working directory
WORKDIR /go/src/app
# add source code
ADD main main
# run main.go
CMD ["go", "run", "main/welcome.go"]