FROM golang:1.19

RUN mkdir /dockUser

ADD . /dockUser/

WORKDIR /dockUser

EXPOSE 8080

RUN go build -o main ./cmd/myapp

CMD ["/dockUser/main"]