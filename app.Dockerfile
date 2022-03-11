FROM golang:1.17-alpine

COPY . /src
WORKDIR /src

RUN go build .

CMD [ "./docker_task" ]