FROM golang:1.20.6
WORKDIR /chatroom
ADD . /chatroom
RUN go mod download
RUN go build -o main .
EXPOSE 8080
CMD ["/chatroom/main"]