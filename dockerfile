FROM golang:1.18

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go build -o deploy-app

CMD ./deploy-app