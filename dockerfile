# build stage

FROM golang:1.14.0-alpine3.11 AS builder

RUN go env -w GOPROXY=https://goproxy.io,direct


# ARG SRC
# RUN echo "src-->"${SRC}

RUN  mkdir /opt2
COPY ./ /opt2/

RUN cd /opt2 && go build -o myapp  /opt2

# final stage
FROM golang:alpine
EXPOSE 8080
RUN  mkdir /opt3
WORKDIR /opt3
COPY --from=builder /opt2 .
CMD ["/opt3/myapp",""]


#FROM golang:1.14.0-alpine3.11

# RUN apk update
# RUN apk add git
#WORKDIR /opt2
#COPY entrypoint.sh  /

#ENTRYPOINT ["/entrypoint.sh"]

#CMD go run /opt2/main.go