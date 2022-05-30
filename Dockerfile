FROM golang:alpine

RUN mkdir /build 

RUN export GO111MODULE=on

#Version github
#RUN cd /build && git clone https://github.com/Shiyro/vade-backend-techtest.git

#Version locale
COPY . /build/vade-backend-techtest

RUN cd /build/vade-backend-techtest/REST-API && go get && go build

EXPOSE 8080

ENTRYPOINT [ "/build/vade-backend-techtest/REST-API/rest-api-service" ]
