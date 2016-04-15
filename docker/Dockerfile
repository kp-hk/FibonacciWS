FROM golang:latest
MAINTAINER Kam Pang <kam_k_pang@yahoo.com>
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o fibon_ws .
EXPOSE 8080
CMD ["/app/fibon_ws"]
