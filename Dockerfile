FROM golang:1.17-alpine

WORKDIR /app

COPY . /app


RUN go build -o /rest-service

EXPOSE 8080

CMD [ "/rest-service"]