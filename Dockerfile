FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN chmod +x start.sh
RUN go get
RUN go build

EXPOSE 80

ENTRYPOINT ./start.sh