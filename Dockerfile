FROM golang:1.9 

ENV GOPATH=/app
ENV WORKPATH=/app/src

COPY src $WORKPATH

WORKDIR $WORKPATH

RUN go get .

RUN go build -o application .

EXPOSE 8080

CMD ["./application"]