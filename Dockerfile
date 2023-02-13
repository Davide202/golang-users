FROM golang:1.19



WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./

COPY src .

#WORKDIR /usr/src/app/src

RUN go mod download && go mod verify


RUN go build -v 
#RUN go build -v -o /usr/local/bin/app ./...
#RUN go build -v -o $WORKDIR ./...

ENV PATH=/usr/src/app
EXPOSE 8080

CMD ["golang-users"]


#You can then build and run the Docker image:

#$ docker build -t users-image:1 .
#$ docker run -it -p 8080:8080 -e mysql_users_host=172.17.0.1 --rm --name users-container users-image:1 